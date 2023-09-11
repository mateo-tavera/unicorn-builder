package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mateo-tavera/unicorn-builder/repository"
	"github.com/mateo-tavera/unicorn-builder/service"
)

func GetUnicornHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Processing new request...")
	// Prepare values for service
	values := r.URL.Query()
	amount, _ := strconv.Atoi(values.Get("amount"))
	if amount <= 0 {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}
	// Input
	names := repository.GetNames()
	adjectives := repository.GetAdjectives()
	// Create service request
	unicorns := service.GetUnicorns(amount, names, adjectives)
	// Map service response
	response, err := json.Marshal(unicorns)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	//Output
	fmt.Println("Request has been completed.")
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
