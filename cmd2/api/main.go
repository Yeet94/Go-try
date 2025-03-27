package main

//The internal/ directory is a Go convention for organizing code that should only be used within your project.
// Any code inside internal/ cannot be imported by other projects,
// ensuring encapsulation and preventing accidental misuse.

//A router maps incoming HTTP requests to specific handler functions. For example:

import (
	"fmt"
	"net/http"

	"github.com/Yeet94/Go_try/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)        // Enables detailed logging with file and line numbers
	var r *chi.Mux = chi.NewRouter() // Creates a new router using the Chi package
	handlers.Handler(r)              // Registers routes and handlers

	fmt.Println("Starting GO Api...")
	fmt.Println("TEST: Listening on port 8000")

	err := http.ListenAndServe("localhost:8000", r) // Starts the server on port 8000
	if err != nil {
		log.Error(err)
	}
}
