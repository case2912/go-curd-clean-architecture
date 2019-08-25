package presenter

import "github.com/case2912/go-curd-clean-architecture/interface/adapter"

type UserCreatePresenter struct {
}

func NewUserCreatePresenter() adapter.UserCreatePresenter {
	return &UserCreatePresenter{}
}
