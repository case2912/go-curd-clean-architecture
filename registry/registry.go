package registry

import (
	"github.com/case2912/go-curd-clean-architecture/adapter"
	"github.com/case2912/go-curd-clean-architecture/application/controller"
	"github.com/case2912/go-curd-clean-architecture/application/interactor"
	"github.com/case2912/go-curd-clean-architecture/application/usecase"
	"github.com/case2912/go-curd-clean-architecture/driver/database"
	"github.com/case2912/go-curd-clean-architecture/driver/presenter"
	"github.com/case2912/go-curd-clean-architecture/driver/repository"
)

type Registry struct {
	sqlConfig *database.SQLConfig
}

func NewRegistry() *Registry {
	return &Registry{}
}
func (r *Registry) SetConfig(sqlConfig *database.SQLConfig) {
	r.sqlConfig = sqlConfig
}
func (r *Registry) NewSQLHandler() adapter.SQLHandler {
	return database.NewSQLHandler(r.sqlConfig)
}

func (r *Registry) NewUserRepository() adapter.UserRepository {
	return repository.NewUserRepository(r.NewSQLHandler())
}

func (r *Registry) NewUserCreatePresenter() adapter.UserCreatePresenter {
	return presenter.NewUserCreatePresenter()
}

func (r *Registry) NewUserCreateInteractor() usecase.UserCreateUsecase {
	return interactor.NewUserCreateInteractor(r.NewUserRepository(), r.NewUserCreatePresenter())
}

func (r *Registry) NewUserController() *controller.UserController {
	return controller.NewUserController(r.NewUserCreateInteractor())
}
