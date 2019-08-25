package presenter

import "github.com/case2912/go-curd-clean-architecture/adapter"

type UserCreatePresenter struct {
}

func NewUserCreatePresenter() adapter.UserCreatePresenter {
	return &UserCreatePresenter{}
}
