package user

import (
	"Test-Task2/internal/entity"
	"context"
)

func (uc *Usecase) GetAllUsers(ctx context.Context, pag int) (*[]*entity.User, error) {
	users, err := uc.user.Users(ctx, pag)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uc *Usecase) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	user, err := uc.user.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *Usecase) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	user, err := uc.user.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *Usecase) CreateUser(ctx context.Context, dto *entity.User) (string, error) {
	err := uc.user.CreateUser(ctx, dto)
	if err != nil {
		return "", err
	}

	user, err := uc.user.GetUserByUsername(ctx, dto.Username)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}

func (uc *Usecase) UpdateUser(ctx context.Context, dto *entity.User) error {
	err := uc.user.Update(ctx, dto)
	if err != nil {
		return err
	}

	return nil
}

func (uc *Usecase) DeleteUser(ctx context.Context, id string) error {
	err := uc.user.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
