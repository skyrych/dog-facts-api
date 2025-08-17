package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	app "github.com/skyrych/dog-facts-api/internal/app/dogfacts"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}
	needyRandomFact := app.NewFactServer()
	server := app.StartServer(port, needyRandomFact)

	serverErrors := make(chan error, 1)

	// Launch the server in a goroutine.
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErrors <- err
		}
	}()

	// --- SHUTDOWN ORCHESTRATION ---
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Server is running on :80. Press Ctrl+C to shut down gracefully.")

	select {
	case err := <-serverErrors:
		log.Fatalf("Fatal server error: %v", err)
	case <-sigs:
		log.Println("Received OS signal. Initiating graceful shutdown...")

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatalf("Server shutdown failed: %v", err)
		}
		log.Println("Server shut down gracefully. Exiting.")
	}
}
