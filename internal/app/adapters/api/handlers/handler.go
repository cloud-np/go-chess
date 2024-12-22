package handlers

import (
	"fmt"
	"net/http"
)

func PlayMove(w http.ResponseWriter, r *http.Request) {
	// Handle non allowed methods
	if r.Method != http.MethodPost {
		http.Error(w, "Ivalid request method", http.StatusMethodNotAllowed)
		return
	}

	fmt.Print(r.Body)
}

func FourOFour(w http.ResponseWriter, r *http.Request) {
	// Handle non allowed methods
	fmt.Print(r.Body)
	http.Error(w, "Ivalid request method", http.StatusMethodNotAllowed)
}

func GetFen(w http.ResponseWriter, r *http.Request) {
	// Handle non allowed methods
	if r.Method != http.MethodGet {
		http.Error(w, "Ivalid request method", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Requested: " + r.URL.Path))
	fmt.Print(r.Body)
}
