package auth

import (
	"Test-Task2/internal/entity"
	"context"
	"errors"
)

func (r *AuthStorageRepo) One(ctx context.Context, username, password string) (*entity.User, error) {
	user, err := r.inMemory.GetUserByUsername(ctx, username)
	if err != nil {
		r.log.Error("error in get user by username")
		return nil, err
	}

	if user.IsAdmin == false {
		r.log.Error("your is not admin")
		return nil, errors.New("your is not admin")
	}

	if user.Password != password {
		r.log.Error("your password is not equal to password in app")
		return nil, errors.New("your password is not equal to password in app")
	}

	return user, nil
}

func (r *AuthStorageRepo) CreateUser(ctx context.Context, user *entity.User) error {
	err := r.inMemory.Set(ctx, user)
	if err != nil {
		r.log.Error("error in create user")
		return err
	}
	return nil
}
