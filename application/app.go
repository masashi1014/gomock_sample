package application

import (
	"github.com/masashi1014/gomock_sample/repository"
)

type App struct {
	UserRepository repository.UserRepository
}

func NewApplication() App{
	return App{
		UserRepository: repository.NewUserRepository(),
	}
}

func(app *App) FindOrCreateUser(id int64) error {
	user := app.UserRepository.GetUserByID(id)

	if !user.IsRegistered {
		_, err := app.UserRepository.Create()
		if err != nil {
			return err
		}
	}
	return nil
}