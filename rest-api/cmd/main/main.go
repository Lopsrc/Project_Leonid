package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"rest-api/m/rest-api/internal/config"
	auth2 "rest-api/m/rest-api/internal/auth"
	auth "rest-api/m/rest-api/internal/auth/db"
	user2 "rest-api/m/rest-api/internal/user"
	user "rest-api/m/rest-api/internal/user/db"
	"rest-api/m/rest-api/pkg/client/postgresql"
	"time"

	"github.com/julienschmidt/httprouter"
)


const(
	pathConfig 			= "rest-api/config/my_local.yaml"
	envLocal   			= "local"
	envDev     			= "dev"
	envProd    			= "prod"
)

func main() {
	cfg := config.GetConfig(pathConfig)

	log := setupLogger(cfg.Env)

	router := httprouter.New()

	postgreSQLClient, err := postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		log.Error("%v", err)
	}

	repositoryUserAuth := auth.NewRepository(postgreSQLClient, log)
	repositoryUserData := user.NewRepository(postgreSQLClient, log)

	log.Info("register user handler")
	authHandler := auth2.NewHandler(repositoryUserAuth, log)
	authHandler.Register(router)
	
	userHandler := user2.NewHandler(repositoryUserData, log)
	userHandler.Register(router)
	start(router, cfg, log)
}

func start(router *httprouter.Router, cfg *config.Config, log *slog.Logger) {
	log.Info("start application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		log.Info("detect app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			panic(err)
		}
		log.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")

		log.Info("listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		log.Info("server is listening unix socket: %s", socketPath)
	} else {
		log.Info("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		log.Info("server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	if listenErr != nil {
		log.Error("%s", listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Error("%s", server.Serve(listener))
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}