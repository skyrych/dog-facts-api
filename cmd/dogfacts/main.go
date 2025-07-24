package main

import (
	"fmt"

	app "github.com/skyrych/dog-facts-api/internal/app/dogfacts"
)

func main() {
	needyRandomFact := app.NewFactServer()
	err := app.StarServer(":8080", needyRandomFact)
	if err != nil {
		fmt.Println("Failed to run a server")
	}
}
