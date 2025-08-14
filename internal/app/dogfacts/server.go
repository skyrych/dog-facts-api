package dogfacts

import (
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type FactServer struct {
	Facts  []string
	Rand   *rand.Rand
	Logged *slog.Logger
}

func NewFactServer() *FactServer {
	return &FactServer{
		Facts: []string{"Needy is great.", "Needy is bright.", "Needy hates cats.", "Needy is overprotective over his food.",
			"Needy loves his yummies.", "Needy is the best hunter.", "Needy loves being with his humans.",
			"Needy has friends and enemies among the dogs.", "Needy naps throughout the day.", "Needy becomes anxious and restless when he wants to pee.",
			"Needy loves sweets.", "Needy hunts cats, mice and frogs", "Needy loves going out with his humans.",
			"Needy prefers local outings to coffee shops over long-distance travel.", "Needy has a main and a spare human."},
		Rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
		Logged: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
}

func StartServer(port string, fs *FactServer) error {
	fs.Logged.Info("Starting Dog Facts API server", "port", port)
	http.HandleFunc("/facts", fs.factsHandler)
	return http.ListenAndServe(port, nil)
}

func (fs *FactServer) factsHandler(res http.ResponseWriter, req *http.Request) {
	fs.Logged.Info("Received a request:", "Method", req.Method, "Path", req.URL.Path)
	if len(fs.Facts) == 0 {
		http.Error(res, "Looks like we are out of facts", 503)
		return
	}
	num := fs.Rand.Intn(len(fs.Facts))
	fmt.Fprint(res, fs.Facts[num], "\n")
}
