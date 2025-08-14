// cmd/dogfacts/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	app "github.com/skyrych/dog-facts-api/internal/app/dogfacts"
)

func main() {

	needyRandomFact := app.NewFactServer()
	err := app.StartServer(":80", needyRandomFact)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-sigs:
		fmt.Printf("Received the signal %s: ", s)
		cancel()
	case <-time.After(10 * time.Second):
		fmt.Println("Graceful shutdown initiated due to timeout.")
		cancel()
	}
}
