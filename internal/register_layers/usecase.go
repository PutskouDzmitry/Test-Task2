package register_layers

import (
	"Test-Task2/internal/usecase/auth"
	"Test-Task2/internal/usecase/user"
)

type GUsecase struct {
	user *user.Usecase
	auth *auth.Usecase
}

func NewGUsecase(repo *GRepository) *GUsecase {
	return &GUsecase{
		user: user.NewUserUsecase(repo.user),
		auth: auth.NewAuthUsecase(repo.auth),
	}
}
