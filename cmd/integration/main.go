package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	authCommon "github.com/erupshis/golang-integration-developer-test/internal/common/auth"
	"github.com/erupshis/golang-integration-developer-test/internal/common/auth/authgrpc"
	authPostgres "github.com/erupshis/golang-integration-developer-test/internal/common/auth/storage/postgres"
	"github.com/erupshis/golang-integration-developer-test/internal/common/db"
	"github.com/erupshis/golang-integration-developer-test/internal/common/hasher"
	"github.com/erupshis/golang-integration-developer-test/internal/common/jwtgenerator"
	"github.com/erupshis/golang-integration-developer-test/internal/common/logger"
	"github.com/erupshis/golang-integration-developer-test/internal/common/utils/deferutils"
	"github.com/erupshis/golang-integration-developer-test/internal/integration/auth"
	"github.com/erupshis/golang-integration-developer-test/internal/integration/config"
	"github.com/erupshis/golang-integration-developer-test/internal/integration/integr"
	"github.com/erupshis/golang-integration-developer-test/internal/integration/server"
	"github.com/erupshis/golang-integration-developer-test/internal/service/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	migrationsFolder = "file://db/migrations/"
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

	ctxWithCancel, cancel := context.WithCancel(context.Background())
	defer cancel()

	// auth.
	dbConfig := db.Config{
		DSN:              cfg.DatabaseDSN,
		MigrationsFolder: migrationsFolder,
	}
	databaseConn, err := db.NewConnection(ctxWithCancel, dbConfig)
	if err != nil {
		logs.Fatalf("failed to connect to users database: %v", err)
	}

	jwtGenerator, err := jwtgenerator.NewJWTGenerator(cfg.JWT, 2)
	hash := hasher.CreateHasher(cfg.HashKey, hasher.TypeSHA256, logs)
	authStorage := authPostgres.NewPostgres(databaseConn, logs)
	authManagerConfig := &authCommon.Config{
		Storage: authStorage,
		JWT:     jwtGenerator,
		Hasher:  hash,
	}
	authManager := authCommon.NewManager(authManagerConfig)
	authController := auth.NewController(authManager)

	// handlers integr.
	defClient := client.NewDefault(cfg.PlayersHost)
	integrationController := integr.NewController(defClient)

	// gRPC server options.
	var opts []grpc.ServerOption
	opts = append(opts, grpc.Creds(insecure.NewCredentials()))
	opts = append(opts, grpc.ChainUnaryInterceptor(
		logger.UnaryServer(logs),
		authgrpc.UnaryServer(jwtGenerator, logs),
	))
	// gRPC server
	srv := server.NewServer(integrationController, authController, "grpc", opts...)
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
