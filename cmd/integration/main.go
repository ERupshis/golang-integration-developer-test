package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/erupshis/golang-integration-developer-test/internal/common/logger"
	"github.com/erupshis/golang-integration-developer-test/internal/common/utils/deferutils"
	"github.com/erupshis/golang-integration-developer-test/internal/integration/config"
	"github.com/erupshis/golang-integration-developer-test/internal/integration/controller"
	"github.com/erupshis/golang-integration-developer-test/internal/integration/interceptors/logging"
	"github.com/erupshis/golang-integration-developer-test/internal/integration/server"
	"github.com/erupshis/golang-integration-developer-test/internal/service/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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

	// handlers controller.
	defClient := client.NewDefault(cfg.PlayersHost)
	grpcController := controller.NewController(defClient)

	// gRPC server options.
	var opts []grpc.ServerOption
	opts = append(opts, grpc.Creds(insecure.NewCredentials()))
	opts = append(opts, grpc.ChainUnaryInterceptor(logging.UnaryServer(logs)))
	// gRPC server
	srv := server.NewServer(grpcController, "grpc", opts...)
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
