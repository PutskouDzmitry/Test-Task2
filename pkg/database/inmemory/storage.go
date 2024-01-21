package inmemory

import (
	"Test-Task2/internal/entity"
	"context"
)

type Storage interface {
	Set(ctx context.Context, dto *entity.User) error
	GetUserById(ctx context.Context, key string) (*entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	GetAll(ctx context.Context, pag int) (*[]*entity.User, error)
	DeleteUser(ctx context.Context, id string) error
	UpdateUser(ctx context.Context, dto *entity.User) error
}
