package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	serviceProvider "github.com/DarYur13/learn-control/internal/app"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	app, err := serviceProvider.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to create app: %s", err.Error())
	}

	go func() {
		if err := app.Run(ctx); err != nil {
			log.Printf("app exited with error: %s", err)
			stop()
		}
	}()

	<-ctx.Done()

	log.Println("Shutting down app...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	app.Shutdown(shutdownCtx)

	log.Println("Graceful shutdown complete.")
}
