package main

import (
	"fmt"
	"go/backend/configs"
	"go/backend/internal/auth"
	"go/backend/internal/link"
	"go/backend/pkg/db"
	"net/http"
)

// 10.1
func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	//Repositories
	linkRepository := link.NewLinkRepository(db)

	// Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("server is start")
	server.ListenAndServe()
}
