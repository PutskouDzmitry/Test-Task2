package auth

import (
	"Test-Task2/internal/entity"
	"Test-Task2/internal/usecase/auth"
	"Test-Task2/pkg/api/jwt"
	ssov1 "Test-Task2/proto"
	"context"
)

type Delivery interface {
	SignIn(ctx context.Context, username, password string) (map[string]interface{}, error)
	SingUp(ctx context.Context, user *entity.User) (string, error)
}

type ServerAuth struct {
	ssov1.UnimplementedAuthServer

	jwt  *jwt.JWT
	auth Delivery
}

func Register(jwt *jwt.JWT, auth *auth.Usecase) *ServerAuth {
	return &ServerAuth{
		jwt:  jwt,
		auth: auth,
	}
}
