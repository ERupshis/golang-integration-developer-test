package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/erupshis/golang-integration-developer-test/cmd/players/docs"
	"github.com/erupshis/golang-integration-developer-test/internal/common/logger"
	"github.com/erupshis/golang-integration-developer-test/internal/common/utils/deferutils"
	"github.com/erupshis/golang-integration-developer-test/internal/players/config"
	"github.com/erupshis/golang-integration-developer-test/internal/players/controller"
	"github.com/erupshis/golang-integration-developer-test/internal/players/models"
	"github.com/erupshis/golang-integration-developer-test/internal/players/server"
	"github.com/erupshis/golang-integration-developer-test/internal/players/storage/inmem"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

var users = map[int64]models.UserDataP{
	1: {1, 1000},
	2: {2, 100},
	3: {3, 0},
}

// @title Players service Swagger API
// @version 1.0
// @description Swagger API for players storage.
// @termsOfService http://swagger.io/terms/

// @contact.name erupshis
// @contact.email e.rupshis@gmail.com

// @BasePath /api/v1
func main() {
	logs, err := logger.NewZap("info")
	if err != nil {
		log.Fatalf("create zap logs: %v", err)
	}
	defer deferutils.ExecSilent(logs.Sync)

	cfg, err := config.Parse()
	if err != nil {
		logs.Fatalf("parse config: %v", err)
	}

	// in-memory users storage.
	userStorage := inmem.NewUserStorage(users)

	// handlers controller.
	httpController := controller.NewController(userStorage, logs)

	// attaching handlers.
	router := chi.NewRouter()
	router.Mount("/api/v1", httpController.Route())
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost%s/swagger/doc.json", cfg.Host)),
	))

	// dataprovider launch.
	srv := server.NewServer(cfg.Host, router, "http")
	srv.Host(cfg.Host)

	go func() {
		listener, err := net.Listen("tcp", cfg.Host)
		if err != nil {
			logs.Fatalf("failed to listen for %s dataprovider: %v", srv.GetInfo(), err)
		}

		if err = srv.Serve(listener); err != nil {
			logs.Infof("http://%s dataprovider refused to start or stop with error: %v", srv.GetInfo(), err)
			return
		}
	}()

	ctxWithCancel, cancel := context.WithCancel(context.Background())
	defer cancel()

	// shutdown.
	idleConnsClosed := make(chan struct{})
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sigCh

		if err = srv.GracefulStop(ctxWithCancel); err != nil {
			logs.Infof("%s dataprovider graceful stop error: %v", srv.GetInfo(), err)
		}

		cancel()
		close(idleConnsClosed)
	}()

	<-idleConnsClosed
	logs.Infof("players service shutdown gracefully")
}
