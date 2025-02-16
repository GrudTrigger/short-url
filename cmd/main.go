package main

import (
	"fmt"
	"go/backend/configs"
	"go/backend/internal/auth"
	"go/backend/internal/link"
	"go/backend/internal/user"
	"go/backend/pkg/db"
	"go/backend/pkg/middleware"
	"net/http"
)

// 12.6
func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	//Repositories
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)

	//Services
	authService := auth.NewAuthService(userRepository)

	// Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		Config:         conf,
	})

	//Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}
	fmt.Println("server is start")
	server.ListenAndServe()
}
