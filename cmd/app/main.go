package main

import (
	"log"
	"net/http"

	"github.com/Wammero/PVZ-service/internal/cache"
	"github.com/Wammero/PVZ-service/internal/config.go"
	"github.com/Wammero/PVZ-service/internal/handler"
	"github.com/Wammero/PVZ-service/internal/repository"
	"github.com/Wammero/PVZ-service/internal/router"
	"github.com/Wammero/PVZ-service/internal/service"
	"github.com/Wammero/PVZ-service/pkg/jwt"
	"github.com/Wammero/PVZ-service/pkg/migration"
)

func main() {
	cfg := config.NewConfig()

	jwt.SetSecret(cfg.JWT.SecretKey)

	connstr := cfg.Database.GetConnStr()
	repo, err := repository.New(connstr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer repo.Close()
	migration.ApplyMigrations(connstr)

	redisClient, err := cache.NewRedisClient(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Password, cfg.Redis.DB)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer func() {
		if err := redisClient.Close(); err != nil {
			log.Printf("Error closing Redis connection: %v", err)
		}
	}()

	svc := service.New(repo, redisClient)
	r := router.New()
	h := handler.New(svc)

	h.SetupRoutes(r)

	log.Printf("Server is running on port %s", cfg.Server.Port)
	if err := http.ListenAndServe(":"+cfg.Server.Port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
