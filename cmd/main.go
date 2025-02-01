package main

import (
	"fmt"
	"go/backend/configs"
	"go/backend/internal/auth"
	"net/http"
)

// 6.8
func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()

	// Handler auth
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("server is start")
	server.ListenAndServe()
}
