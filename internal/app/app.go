package app

import (
	"context"
	"github.com/MuZaZaVr/account-service/internal/config"
	"github.com/MuZaZaVr/account-service/internal/handler"
	"github.com/MuZaZaVr/account-service/internal/handler/swagger"
	"github.com/MuZaZaVr/account-service/internal/repository"
	"github.com/MuZaZaVr/account-service/internal/server"
	"github.com/MuZaZaVr/account-service/internal/service"
	"github.com/MuZaZaVr/account-service/pkg/auth"
	"github.com/MuZaZaVr/account-service/pkg/database/mongo"
	"log"
	"os"
	"os/signal"
	"time"
)

func Run(configPath string) {
	ctx := context.Background()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	/* Config */
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatalf("Init config error: %s", err)
	}
	log.Printf("Initialize config: %v", cfg)

	/* Mongo */
	db, err := mongo.NewMongo(ctx, cfg.Mongo)
	if err != nil {
		log.Fatalf("Init db error: %s", err)
	}
	log.Printf("Initialize database: %v", db.Name())

	/* Repositories */
	repos := repository.NewRepositories(db)
	log.Printf("Initialize repositories: %v", repos)

	/* JWT manager */
	tokenManager, err := auth.NewManager(cfg.JWT.SigningKey)
	if err != nil {
		log.Fatalf("Init jwt error: %v", err)
	}
	log.Printf("Initialize JWT manager: %v", tokenManager)

	/* Services */
	services := service.NewServices(service.Deps{
		Repos:        repos,
		TokenManager: tokenManager,
	})
	log.Printf("Initialize services: %v", services)

	/* Handler */
	newHandler := handler.NewHandler(services, tokenManager)

	some := newHandler.GetRoute("companyRouter").GetHandler()
	log.Printf("Initialize handler: %v", some)

	swagger.HandlerSwagger(newHandler)

	/* Server */
	newServer := server.NewServer(cfg, newHandler)
	log.Printf("Initialize server: %v", newServer)

	log.Printf("Starting service: %v", newServer)
	go startService(ctx, newServer)

	<-stop

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := newServer.Stop(ctx); err != nil {
		log.Fatalf("Failed to stop the newServer: %s", err.Error())
	}

	log.Printf("Shutting down server...")
}

func startService(ctx context.Context, server *server.Server) {
	if err := server.Run(); err != nil {
		log.Fatal(ctx, "Server shutdown: ", err.Error())
	}
}