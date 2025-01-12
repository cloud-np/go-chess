package handlers

import (
	"fmt"
	"net/http"
)
func FourOFour(w http.ResponseWriter, r *http.Request) {
	// Handle non allowed methods
	fmt.Print(r.Body)
	http.Error(w, "Ivalid request method", http.StatusMethodNotAllowed)
}

