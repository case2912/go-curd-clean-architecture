package controller

import (
	"github.com/case2912/go-curd-clean-architecture/application/usecase"
	"github.com/case2912/go-curd-clean-architecture/domain"
)

type UserController struct {
	UserCreateUsecase usecase.UserCreateUsecase
}

func NewUserController(userCreateUsecase usecase.UserCreateUsecase) *UserController {
	return &UserController{
		UserCreateUsecase: userCreateUsecase,
	}
}

func (controller *UserController) CreateUser(user domain.User) domain.User {
	return controller.UserCreateUsecase.Handle(user)
}
