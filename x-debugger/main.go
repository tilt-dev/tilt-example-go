package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ExampleRouter struct {
	*mux.Router
}

func NewExampleRouter() *ExampleRouter {
	r := mux.NewRouter()

	r.Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		if err := enc.Encode(map[string]interface{}{"status": "healthy"}); err != nil {
			panic(err)
		}
	})

	return &ExampleRouter{
		Router: r,
	}
}

func main() {
	http.Handle("/", NewExampleRouter())

	log.Println("Serving on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Server exited with: %v", err)
	}
}
