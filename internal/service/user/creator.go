package user

import (
	"shop-smart-api/internal/controller/http/types"
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/infrastructure/repository"
)

type Creator interface {
	Create(phone string, chanel *types.Channel) (*entity.User, error)
}

type creator struct {
	repository repository.UserRepository
}

func CreateCreator(r repository.UserRepository) Creator {
	return &creator{r}
}

func (c *creator) Create(resource string, chanel *types.Channel) (*entity.User, error) {
	if chanel.IsEmail() {
		return c.repository.Store("", resource, "", []entity.Role{entity.UserRole})
	}
	if chanel.IsCode() {
		return c.repository.Store("", "", resource, []entity.Role{entity.UserRole})
	}
	return c.repository.Store(resource, "", "", []entity.Role{entity.UserRole})
}
