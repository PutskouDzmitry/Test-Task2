package user

import (
	"Test-Task2/pkg/database/inmemory"
	"go.uber.org/zap"
)

type UserStorageRepo struct {
	inMemory inmemory.Storage
	log      *zap.Logger
}

func UserStorage(memory inmemory.Storage, log *zap.Logger) *UserStorageRepo {
	return &UserStorageRepo{
		inMemory: memory,
		log:      log,
	}
}
