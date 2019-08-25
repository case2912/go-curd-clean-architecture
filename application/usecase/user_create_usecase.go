package usecase

import "github.com/case2912/go-curd-clean-architecture/domain"

type UserCreateUsecase interface {
	Handle(domain.User) domain.User
}
