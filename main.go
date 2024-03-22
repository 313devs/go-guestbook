package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

)

func main() {

	router:= chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/ping", basicHandler)

	server := &http.Server{
		Addr:    ":3001",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to listen to server:", err)
	}
}
func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
