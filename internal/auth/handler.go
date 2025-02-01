package auth

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go/backend/configs"
	"go/backend/pkg/response"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//Читаем body
		var payload LoginRequest
		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			response.Json(w, err.Error(), 402)
			return
		}

		validate := validator.New()
		err = validate.Struct(payload)
		if err != nil {
			response.Json(w, err.Error(), 402)
			return
		}

		data := LoginResponse{
			Token: "123",
		}

		response.Json(w, data, 200)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Register")
	}
}
