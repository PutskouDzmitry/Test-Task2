package auth

import (
	"Test-Task2/internal/entity"
	"context"
	"errors"
)

func (uc *Usecase) SignIn(ctx context.Context, username, password string) (map[string]interface{}, error) {
	account, err := uc.auth.One(ctx, username, password)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, errors.New("app doesn't find your in database")
	}

	return map[string]interface{}{
		"user_id": account.ID,
		"role":    "admin",
	}, nil
}

func (uc *Usecase) SingUp(ctx context.Context, user *entity.User) (string, error) {
	err := uc.auth.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}

	account, err := uc.auth.One(ctx, user.Username, user.Password)
	if err != nil {
		return "", err
	}

	return account.ID, nil
}
