package dogfacts

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type FactServer struct {
	Facts []string
	Rand  *rand.Rand
}

func NewFactServer() *FactServer {
	return &FactServer{
		Facts: []string{"Needy is great", "Needy is bright", "Needy hates cats", "Needy is overprotective over his food",
			"Needy loves his yummies", "Needy is the best hunter", "Needy loves being with his humans.",
			"Needy has friends and enemies among the dogs.", "Needy naps throughout the day.", "Needy becomes anxious and restless when he wants to pee.",
			"Needy loves sweets.", "Needy hunts cats, mice and frogs", "Needy loves going out with his humans.",
			"Needy prefers local outings to coffee shops over long-distance travel."},
		Rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func StartServer(port string, cert string, key string, fs *FactServer) error {
	http.HandleFunc("/facts", fs.factsHandler)
	return http.ListenAndServeTLS(port, cert, key, nil)
}

func (fs *FactServer) factsHandler(res http.ResponseWriter, req *http.Request) {
	num := fs.Rand.Intn(len(fs.Facts))
	fmt.Fprint(res, fs.Facts[num], "\n")
}
