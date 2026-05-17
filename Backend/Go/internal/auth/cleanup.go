package auth

import (
	"context"
	"log"
	"time"
)

func StartOTPCleanup(ctx context.Context, repo *Repository) {
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for {
			select {
			case <-ticker.C:
				if err := repo.DeleteExpiredOTPs(ctx); err != nil {
					log.Printf("otp cleanup error: %v", err)
				}
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}
	}()
}
