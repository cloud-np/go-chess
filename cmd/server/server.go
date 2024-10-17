package server

import (
	"fmt"
	"gochess/internal/handlers"
	"net/http"
)

func NewServer(port string) {
	// http.HandleFunc("/fen", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	// })
	// Add handlers
	// TODO: Later on add aunthetication/authorization middlewares etc
	http.HandleFunc("/fen", handlers.GetFen)
	http.HandleFunc("/*", handlers.Ep)

	fmt.Println("Server starting on " + port + "...")

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
