package adapter

import "github.com/case2912/go-curd-clean-architecture/domain"

type UserController interface {
	CreateUser(domain.User) domain.User
}
