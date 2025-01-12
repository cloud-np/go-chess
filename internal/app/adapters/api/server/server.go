package server

import (
	"fmt"
	"gochess/internal/app/adapters/api/handlers"
	"gochess/internal/app/adapters/api/middleware"
	"net/http"
)

func RunServer(port string) {

	http.HandleFunc("/fen", middleware.ChainMiddlewares(handlers.SetFen,
		middleware.CorrectMethods([]string{http.MethodPost}),
		middleware.CorsMiddleware,
		middleware.JSONMiddleware,
	))
	http.HandleFunc("/*", middleware.CorsMiddleware(handlers.FourOFour))

	fmt.Println("Server starting on " + port + "...")

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
