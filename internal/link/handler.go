package link

import (
	"fmt"
	"net/http"
)

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

type LinkHandler struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(route *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}
	route.HandleFunc("POST /link", handler.Create())
	route.HandleFunc("GET /{hash}", handler.GoTo())
	route.HandleFunc("PATCH /link/{id}", handler.Update())
	route.HandleFunc("DELETE /link/{id}", handler.Delete())
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		fmt.Println("CREATE")
	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		fmt.Println("GET LINK")
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		fmt.Println("UPDATE")
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		// Получение id из query запроса
		id := request.PathValue("id")
		fmt.Println(id)
	}
}
