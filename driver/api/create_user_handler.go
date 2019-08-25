package api

import (
	"encoding/json"
	"net/http"

	"github.com/case2912/go-curd-clean-architecture/domain"
	"github.com/case2912/go-curd-clean-architecture/interface/adapter"
)

type CreateUserHandler struct {
	Controller adapter.UserController
}

func NewCreateUserHandler(controller adapter.UserController) http.Handler {
	return &CreateUserHandler{
		Controller: controller,
	}
}
func (handler *CreateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := handler.Controller.CreateUser(domain.User{
		UserName: "testuser",
	})
	json, err := json.Marshal(user)
	if err != nil {

	}
	w.Header().Set("Content-Type", "app/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
