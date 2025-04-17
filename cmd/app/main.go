package main

import (
	"log"
	"net/http"

	"github.com/Wammero/PVZ-service/internal/config.go"
	"github.com/Wammero/PVZ-service/internal/cron"
	"github.com/Wammero/PVZ-service/internal/handler"
	"github.com/Wammero/PVZ-service/internal/repository"
	"github.com/Wammero/PVZ-service/internal/router"
	"github.com/Wammero/PVZ-service/internal/service"
	"github.com/Wammero/PVZ-service/pkg/jwt"
	"github.com/Wammero/PVZ-service/pkg/migration"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	svc := service.New(repo)
	r := router.New()
	h := handler.New(svc)

	h.SetupRoutes(r)
	r.Handle("/metrics", promhttp.Handler())

	c := cron.New(repo)
	c.Start()
	defer c.Stop()

	log.Printf("Server is running on port %s", cfg.Server.Port)
	if err := http.ListenAndServe(":"+cfg.Server.Port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
