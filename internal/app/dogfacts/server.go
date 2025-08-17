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

// StartServer now returns a shutdown-aware *http.Server instance.
func StartServer(port string, fs *FactServer) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/facts", fs.factsHandler)
	mux.HandleFunc("/healthz", healthzHandler)

	srv := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	fs.Logged.Info("Creating a shutdown-aware HTTP server", "port", port)
	return srv
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

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ok")
}
