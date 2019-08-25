package adapter

import (
	"github.com/case2912/go-curd-clean-architecture/domain"
)

type UserRepository interface {
	Store(domain.User) domain.User
	FindByUserName(domain.User) domain.User
}
