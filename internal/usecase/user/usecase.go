package user

import "Test-Task2/internal/interfaces"

type Usecase struct {
	user interfaces.User
}

func NewUserUsecase(
	user interfaces.User,
) *Usecase {
	return &Usecase{
		user: user,
	}
}
