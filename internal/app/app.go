package app

import (
	"context"
	"fmt"
	"github.com/MuZaZaVr/account-service/internal/config"
	"github.com/MuZaZaVr/account-service/internal/repository"
	"github.com/MuZaZaVr/account-service/internal/service"
	"github.com/MuZaZaVr/account-service/pkg/auth"
	"github.com/MuZaZaVr/account-service/pkg/database/mongo"
	"log"
)

func Run(configPath string) {
	ctx := context.Background()

	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatalf("Error loading config file: %s", err)
	}

	db, err := mongo.NewMongo(ctx, cfg.Mongo)

	repos := repository.NewRepositories(db)

	tokenManager, err := auth.NewManager(cfg.JWT.SigningKey)
	if err != nil {
		log.Fatalf("Error while creating token manager: %v", err)
	}

	services := service.NewServices(service.Deps{
		Repos:        repos,
		TokenManager: tokenManager,
	})

	fmt.Printf("Services: %v", services) // remove after completing
}