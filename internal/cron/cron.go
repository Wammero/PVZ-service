package cron

import (
	"log"

	"github.com/Wammero/PVZ-service/internal/repository"
	"github.com/robfig/cron/v3"
)

type Cron struct {
	c    *cron.Cron
	repo *repository.Repository
}

func New(repo *repository.Repository) *Cron {
	c := cron.New(cron.WithSeconds())
	cr := &Cron{
		c:    c,
		repo: repo,
	}
	log.Println("Registering cron tasks...")
	cr.registerCleanupTask()
	return cr
}

func (cr *Cron) Start() {
	log.Println("Starting cron scheduler...")
	cr.c.Start()
}

func (cr *Cron) Stop() {
	log.Println("Stopping cron scheduler...")
	cr.c.Stop()
}
