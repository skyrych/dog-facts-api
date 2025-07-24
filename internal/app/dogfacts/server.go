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
		Facts: []string{"Needy is great", "Needy is bright", "Needy hates cats"},
		Rand:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func StarServer(port string, fs *FactServer) error {
	http.HandleFunc("/facts", fs.factsHandler)
	return http.ListenAndServe(port, nil)
}

func (fs *FactServer) factsHandler(res http.ResponseWriter, req *http.Request) {
	num := fs.Rand.Intn(len(fs.Facts))
	fmt.Fprint(res, fs.Facts[num])
}
