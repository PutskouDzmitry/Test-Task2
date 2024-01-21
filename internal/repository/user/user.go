package user

import (
	"Test-Task2/internal/entity"
	"context"
)

func (uc *UserStorageRepo) Users(ctx context.Context, pag int) (*[]*entity.User, error) {
	users, err := uc.inMemory.GetAll(ctx, pag)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uc *UserStorageRepo) CreateUser(ctx context.Context, user *entity.User) error {
	err := uc.inMemory.Set(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserStorageRepo) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	user, err := uc.inMemory.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserStorageRepo) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	user, err := uc.inMemory.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserStorageRepo) Update(ctx context.Context, user *entity.User) error {
	err := uc.inMemory.UpdateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserStorageRepo) Delete(ctx context.Context, id string) error {
	err := uc.inMemory.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
