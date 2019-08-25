package repository

import (
	"fmt"

	"github.com/case2912/go-curd-clean-architecture/adapter"
	"github.com/case2912/go-curd-clean-architecture/domain"
)

type UserRepository struct {
	SQLHandler adapter.SQLHandler
}

func NewUserRepository(sqlHandler adapter.SQLHandler) adapter.UserRepository {
	return &UserRepository{
		SQLHandler: sqlHandler,
	}
}

func (repo *UserRepository) Store(user domain.User) domain.User {
	row, err := repo.SQLHandler.Query("INSERT INTO account (user_name) VALUES ($1) RETURNING user_name;", user.UserName)
	defer row.Close()
	if err != nil {
		err = fmt.Errorf("Invalid query error:\n%s", err.Error())
	}
	row.Next()
	var userName string
	row.Scan(&userName)
	user.UserName = userName
	return user
}
func (repo *UserRepository) FindByUserName(user domain.User) domain.User {
	return user
}
