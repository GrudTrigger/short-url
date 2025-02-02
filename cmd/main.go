package main

import (
	"fmt"
	"go/backend/configs"
	"go/backend/internal/auth"
	"go/backend/pkg/db"
	"net/http"
)

// 8.5
func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
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
