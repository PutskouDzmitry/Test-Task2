package interfaces

import (
	"Test-Task2/internal/entity"
	"context"
)

type User interface {
	Users(ctx context.Context, pag int) (*[]*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id string) error
}
