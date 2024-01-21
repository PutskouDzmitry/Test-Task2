package auth

import (
	"Test-Task2/internal/delivery/errors"
	"Test-Task2/internal/entity"
	ssov1 "Test-Task2/proto"
	"context"
	"github.com/google/uuid"
)

func (a ServerAuth) Register(ctx context.Context, reg *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	if reg.GetPassword() == "" {
		return nil, errors.NewErrInternal("your password is empty")
	}
	if reg.GetUsername() == "" {
		return nil, errors.NewErrInternal("your username is empty")
	}
	if reg.GetEmail() == "" {
		return nil, errors.NewErrInternal("your email is empty")
	}

	user := &entity.User{
		Email:    reg.GetEmail(),
		Username: reg.GetUsername(),
		Password: reg.GetPassword(),
		IsAdmin:  true,
	}

	id, err := a.auth.SingUp(ctx, user)
	if err != nil {
		return nil, errors.NewErrInternal(err.Error())
	}

	return &ssov1.RegisterResponse{UserId: id}, nil
}
func (a ServerAuth) Login(ctx context.Context, login *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	if login.GetPassword() == "" {
		return nil, errors.NewErrInternal("your password is empty")
	}
	if login.GetUsername() == "" {
		return nil, errors.NewErrInternal("your username is empty")
	}

	m, err := a.auth.SignIn(ctx, login.GetUsername(), login.GetPassword())
	if err != nil {
		return nil, errors.NewErrInternal(err.Error())
	}

	var userID string
	var role string
	if m == nil {
		userID = uuid.NewString()
		role = "user"
	} else {
		role = m["role"].(string)
		userID = m["user_id"].(string)
	}

	token, err := a.jwt.CreateAccessToken(
		userID,
		string(role),
	)
	if err != nil {
		return nil, errors.NewErrInternal(err.Error())
	}

	return &ssov1.LoginResponse{
		Token: token,
	}, nil
}
func (a ServerAuth) mustEmbedUnimplementedAuthServer() {}
