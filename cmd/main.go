package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mateo-tavera/unicorn-builder/api"
	"github.com/mateo-tavera/unicorn-builder/stack"
)

func main() {
	// Create a store
	store := stack.NewUnicornStack()
	stack.SetUnicornStack(store)

	// Start storing unicorns in the stack
	go stack.CreateBacklogUnicorns(store)

	// Resgiter router to the service
	r := mux.NewRouter()
	r.HandleFunc("/api/get-unicorn", api.GetUnicornHandler).Methods(http.MethodGet)
	http.Handle("/", r)

	// Create the server and start listening for requests
	log.Printf("Listening at port: %d", 8888)
	http.ListenAndServe(":8888", nil)
}
