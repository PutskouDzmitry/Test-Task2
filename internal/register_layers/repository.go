package register_layers

import (
	"Test-Task2/internal/interfaces"
	"Test-Task2/internal/repository/auth"
	"Test-Task2/internal/repository/user"
	"Test-Task2/pkg/database/inmemory"
	"go.uber.org/zap"
)

type GRepository struct {
	user interfaces.User
	auth interfaces.Auth
}

func NewGRepository(
	db inmemory.Storage,
	log *zap.Logger,
) *GRepository {
	return &GRepository{
		user: user.UserStorage(db, log),
		auth: auth.AuthStorage(db, log),
	}
}
