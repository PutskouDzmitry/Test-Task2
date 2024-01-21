package auth

import (
	"Test-Task2/pkg/database/inmemory"
	"go.uber.org/zap"
)

type AuthStorageRepo struct {
	inMemory inmemory.Storage
	log      *zap.Logger
}

func AuthStorage(memory inmemory.Storage, log *zap.Logger) *AuthStorageRepo {
	return &AuthStorageRepo{
		inMemory: memory,
		log:      log,
	}
}
