package interactor

import (
	"github.com/case2912/go-curd-clean-architecture/application/usecase"
	"github.com/case2912/go-curd-clean-architecture/domain"
	"github.com/case2912/go-curd-clean-architecture/interface/adapter"
)

type UserCreateInteractor struct {
	UserRepository      adapter.UserRepository
	UserCreatePresenter adapter.UserCreatePresenter
}

func NewUserCreateInteractor(userRepository adapter.UserRepository, userCreatePresenter adapter.UserCreatePresenter) usecase.UserCreateUsecase {
	return &UserCreateInteractor{
		UserRepository:      userRepository,
		UserCreatePresenter: userCreatePresenter,
	}
}

func (interactor *UserCreateInteractor) Handle(user domain.User) domain.User {
	return interactor.UserRepository.Store(user)
}
