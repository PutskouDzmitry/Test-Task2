package interfaces

import (
	"Test-Task2/internal/entity"
	"context"
)

type Auth interface {
	One(ctx context.Context, username, password string) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
}
