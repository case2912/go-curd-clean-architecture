package api

import (
	"encoding/json"
	"net/http"

	"github.com/case2912/go-curd-clean-architecture/application/controller"
	"github.com/case2912/go-curd-clean-architecture/domain"
)

func Create(w http.ResponseWriter, r *http.Request, controller *controller.UserController) {
	w.Header().Set("Content-Type", "app/json")
	user := controller.CreateUser(domain.User{
		UserName: "aaaa",
	})
	json, err := json.Marshal(user)
	if err != nil {

	}
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
