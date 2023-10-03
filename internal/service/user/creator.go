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

func (c *creator) Create(phone string, chanel *types.Channel) (*entity.User, error) {
	if chanel.IsEmail() {
		return c.repository.Store("", phone, []entity.Role{entity.UserRole})
	}
	return c.repository.Store(phone, "", []entity.Role{entity.UserRole})
}
