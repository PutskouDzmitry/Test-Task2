package inmemory

import (
	"Test-Task2/internal/entity"
	"context"
	"errors"
	"github.com/google/uuid"
	"sync"
)

type InMemory struct {
	sync.RWMutex
	items map[string]entity.UserInMemory
}

func NewInMemory() *InMemory {

	items := make(map[string]entity.UserInMemory)

	cache := InMemory{
		items: items,
	}

	return &cache
}

func (c *InMemory) Set(ctx context.Context, dto *entity.User) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	c.Lock()

	id := uuid.New().String()

	for _, value := range c.items {
		if value.Email == dto.Email {
			return errors.New("your email is consist in database, write new email")
		}

		if value.Username == dto.Username {
			return errors.New("your username is consist in database, write new username")
		}
	}

	defer c.Unlock()

	c.items[id] = entity.UserInMemory{
		Email:    dto.Email,
		Username: dto.Username,
		Password: dto.Password,
		IsAdmin:  dto.IsAdmin,
	}
	return nil
}

func (c *InMemory) GetUserById(ctx context.Context, key string) (*entity.User, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	dto, found := c.items[key]

	if !found {
		return nil, errors.New("rows is empty")
	}

	return &entity.User{
		ID:       key,
		Email:    dto.Email,
		Username: dto.Username,
		Password: dto.Password,
		IsAdmin:  dto.IsAdmin,
	}, nil
}

func (c *InMemory) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	for key, value := range c.items {
		if value.Username == username {
			return &entity.User{
				ID:       key,
				Email:    value.Email,
				Username: value.Username,
				Password: value.Password,
				IsAdmin:  value.IsAdmin,
			}, nil
		}
	}

	return nil, errors.New("rows is empty")
}

func (c *InMemory) GetAll(ctx context.Context, pag int) (*[]*entity.User, error) {

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	var i int
	dto := make([]*entity.User, 0)
	for key, value := range c.items {
		if i == pag {
			return &dto, nil
		}

		dto = append(dto, &entity.User{
			ID:       key,
			Email:    value.Email,
			Username: value.Username,
			Password: value.Password,
			IsAdmin:  value.IsAdmin,
		})

		i++
	}

	return &dto, nil
}

func (c *InMemory) DeleteUser(ctx context.Context, id string) error {

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, found := c.items[id]
	if found == false {
		return errors.New("empty rows")
	}

	c.Lock()
	defer c.Unlock()

	delete(c.items, id)

	return nil
}

func (c *InMemory) UpdateUser(ctx context.Context, dto *entity.User) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, found := c.items[dto.ID]
	if found == false {
		return errors.New("empty rows")
	}

	c.Lock()
	defer c.Unlock()

	c.items[dto.ID] = entity.UserInMemory{
		Email:    dto.Email,
		Username: dto.Username,
		Password: dto.Password,
		IsAdmin:  dto.IsAdmin,
	}

	return nil
}
