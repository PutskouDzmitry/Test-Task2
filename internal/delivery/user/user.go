package user

import (
	"Test-Task2/internal/delivery/errors"
	"Test-Task2/internal/entity"
	"Test-Task2/pkg/utils/api"
	ssov1 "Test-Task2/proto"
	"context"
	"fmt"
)

func (s ServerUser) GetUsers(ctx context.Context, pag *ssov1.ReadAllUsersRequest) (*ssov1.ReadAllUsersResponse, error) {
	var count = 0

	if &pag.Count == nil {
		count = -1
	} else {
		count = int(pag.Count)
	}

	users, err := s.user.GetAllUsers(ctx, count)
	if err != nil {
		return nil, errors.NewErrInternal(err.Error())
	}

	usersResp := make([]*ssov1.User, 0)

	for _, value := range *users {
		usersResp = append(usersResp, &ssov1.User{
			Id:       value.ID,
			Email:    value.Email,
			Username: value.Username,
			Password: value.Password,
			Admin:    value.IsAdmin,
		})
	}

	return &ssov1.ReadAllUsersResponse{
		User: usersResp,
	}, nil
}
func (s ServerUser) GetUserById(ctx context.Context, dto *ssov1.GetUserByIdRequest) (*ssov1.GetUserResponse, error) {
	if dto.GetId() == "" {
		return nil, errors.NewErrInternal("your id is empty")
	}

	user, err := s.user.GetUserByID(ctx, dto.GetId())
	if err != nil {
		return nil, errors.NewErrNotAllowed(err.Error())
	}

	return &ssov1.GetUserResponse{
		User: &ssov1.User{
			Id:       user.ID,
			Email:    user.Email,
			Username: user.Username,
			Password: user.Password,
			Admin:    user.IsAdmin,
		},
	}, nil

}
func (s ServerUser) GetUserByUsername(ctx context.Context, dto *ssov1.GetUserByUsernameRequest) (*ssov1.GetUserResponse, error) {
	if dto.GetUsername() == "" {
		return nil, errors.NewErrInternal("your username is empty")
	}

	user, err := s.user.GetUserByUsername(ctx, dto.GetUsername())
	if err != nil {
		return nil, errors.NewErrInternal(err.Error())
	}

	return &ssov1.GetUserResponse{
		User: &ssov1.User{
			Id:       user.ID,
			Email:    user.Email,
			Username: user.Username,
			Password: user.Password,
			Admin:    user.IsAdmin,
		},
	}, nil

}
func (s ServerUser) NewUser(ctx context.Context, dto *ssov1.AddUserRequest) (*ssov1.AddUserResponse, error) {
	mc, err := api.GetMapClaims(dto.GetToken(), s.jwt)
	if err != nil {
		return nil, errors.NewErrNotAllowed(err.Error())
	}

	role, ok := mc["role"]
	if ok == false {
		return nil, errors.NewErrNotAllowed("you don't have permission")
	}
	if role != "admin" {
		return nil, errors.NewErrNotAllowed("you don't have permission")
	}

	id, err := s.user.CreateUser(ctx, &entity.User{
		Email:    dto.User.GetEmail(),
		Username: dto.User.GetUsername(),
		Password: dto.User.GetUsername(),
		IsAdmin:  dto.User.GetAdmin(),
	})

	if err != nil {
		return nil, errors.NewErrInternal(err.Error())
	}

	return &ssov1.AddUserResponse{
		Id: id,
	}, nil
}
func (s ServerUser) Update(ctx context.Context, dto *ssov1.UpdateUserRequest) (*ssov1.UpdateUserResponse, error) {
	mc, err := api.GetMapClaims(dto.GetToken(), s.jwt)
	if err != nil {
		return nil, errors.NewErrNotAllowed(err.Error())
	}

	role, ok := mc["role"]
	if ok == false {
		return nil, errors.NewErrNotAllowed("you don't have permission")
	}
	if role != "admin" {
		return nil, errors.NewErrNotAllowed("you don't have permission")
	}

	fmt.Println(dto.GetUser())

	err = s.user.UpdateUser(ctx, &entity.User{
		ID:       mc["user_id"].(string),
		Email:    dto.User.GetEmail(),
		Username: dto.User.GetUsername(),
		Password: dto.User.GetPassword(),
		IsAdmin:  dto.User.GetAdmin(),
	})
	if err != nil {
		return nil, errors.NewErrInternal(err.Error())
	}

	return &ssov1.UpdateUserResponse{}, nil
}
func (s ServerUser) DeleteUser(ctx context.Context, dto *ssov1.DeleteUserRequest) (*ssov1.DeleteUserResponse, error) {
	mc, err := api.GetMapClaims(dto.GetToken(), s.jwt)
	if err != nil {
		return nil, errors.NewErrNotAllowed(err.Error())
	}

	role, ok := mc["role"]
	if ok == false {
		return nil, errors.NewErrNotAllowed("you don't have permission")
	}
	if role != "admin" {
		return nil, errors.NewErrNotAllowed("you don't have permission")
	}

	err = s.user.DeleteUser(ctx, mc["user_id"].(string))
	if err != nil {
		return nil, errors.NewErrInternal(err.Error())
	}

	return &ssov1.DeleteUserResponse{}, nil
}
func (ServerUser) mustEmbedUnimplementedUserServiceServer() {}
