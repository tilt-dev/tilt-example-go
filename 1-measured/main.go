package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/windmilleng/tilt-example-go/pkg/start"
)

func calcUpdateDuration() time.Duration {
	if start.StartTime.IsZero() {
		return 0
	}
	return time.Since(start.StartTime)
}

type ExampleRouter struct {
	*mux.Router
	tmpl           *template.Template
	updateDuration time.Duration
}

func NewExampleRouter() (*ExampleRouter, error) {
	r := mux.NewRouter()

	tmpl, err := template.ParseGlob("./web/templates/*.tmpl")
	if err != nil {
		return nil, err
	}

	updateDuration := calcUpdateDuration()
	router := &ExampleRouter{
		Router:         r,
		tmpl:           tmpl,
		updateDuration: updateDuration,
	}

	fs := http.FileServer(http.Dir("./web"))
	r.HandleFunc("/", router.index)
	r.PathPrefix("/").Handler(fs)

	return router, nil
}

func (r *ExampleRouter) updateTimeDisplay() string {
	if r.updateDuration != 0 {
		return r.updateDuration.Truncate(100 * time.Millisecond).String()
	}
	return "N/A"
}

func (r *ExampleRouter) index(w http.ResponseWriter, req *http.Request) {
	err := r.tmpl.ExecuteTemplate(w, "index.tmpl", map[string]string{
		"Duration": r.updateTimeDisplay(),
	})
	if err != nil {
		log.Printf("index: %v")
	}
}

func main() {
	router, err := NewExampleRouter()
	if err != nil {
		log.Fatalf("Router creation failed: %v", err)
	}
	http.Handle("/", router)

	log.Println("Serving on port 8000")
	log.Printf("Deploy time: %s\n", router.updateTimeDisplay())
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Server exited with: %v", err)
	}
}
