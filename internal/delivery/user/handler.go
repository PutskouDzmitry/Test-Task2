package user

import (
	"Test-Task2/internal/entity"
	"Test-Task2/internal/usecase/user"
	"Test-Task2/pkg/api/jwt"
	ssov1 "Test-Task2/proto"
	"context"
)

type Delivery interface {
	GetAllUsers(ctx context.Context, pag int) (*[]*entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	CreateUser(ctx context.Context, dto *entity.User) (string, error)
	UpdateUser(ctx context.Context, dto *entity.User) error
	DeleteUser(ctx context.Context, id string) error
}

type ServerUser struct {
	ssov1.UnimplementedUserServiceServer
	jwt  *jwt.JWT
	user Delivery
}

func Register(jwt *jwt.JWT, user *user.Usecase) *ServerUser {
	return &ServerUser{
		jwt:  jwt,
		user: user,
	}
}
