// cmd/dogfacts/main.go
package main

import (
	"log"

	app "github.com/skyrych/dog-facts-api/internal/app/dogfacts" // <--- ENSURE THIS LINE IS CORRECT!
)

func main() {
	needyRandomFact := app.NewFactServer()
	err := app.StarServer(":8080", needyRandomFact)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
