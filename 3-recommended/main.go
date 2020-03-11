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

func (r *ExampleRouter) index(w http.ResponseWriter, req *http.Request) {
	updateTimeString := "N/A"
	if r.updateDuration != 0 {
		updateTimeString = r.updateDuration.Truncate(100 * time.Millisecond).String()
	}

	congrats := r.updateDuration != 0 && r.updateDuration < 2*time.Second
	err := r.tmpl.ExecuteTemplate(w, "index.tmpl", map[string]interface{}{
		"Duration": updateTimeString,
		"Congrats": congrats,
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
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Server exited with: %v", err)
	}
}
