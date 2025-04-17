package cron

import (
	"context"
	"log"
	"time"
)

func (cr *Cron) registerCleanupTask() {
	_, err := cr.c.AddFunc("0 3 * * * *", func() {
		log.Println("Starting cleanup of inactive products...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		tx, err := cr.repo.Pool().Begin(ctx)
		if err != nil {
			log.Printf("Cleanup failed to start transaction: %v\n", err)
			return
		}
		defer func() {
			if err != nil {
				_ = tx.Rollback(ctx)
			} else {
				_ = tx.Commit(ctx)
			}
		}()

		err = cr.repo.CleanupInactive(ctx, tx)
		if err != nil {
			log.Printf("Cleanup execution failed: %v\n", err)
			return
		}

		log.Println("Cleanup of inactive products completed successfully.")
	})
	if err != nil {
		log.Fatalf("Failed to register cleanup task in cron: %v", err)
	}
}
