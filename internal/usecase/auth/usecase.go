package auth

import "Test-Task2/internal/interfaces"

type Usecase struct {
	auth interfaces.Auth
}

func NewAuthUsecase(
	auth interfaces.Auth,
) *Usecase {
	return &Usecase{
		auth: auth,
	}
}
