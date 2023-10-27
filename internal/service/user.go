package service

import (
	"log"
	"shop-smart-api/internal/controller/http/types"
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/pkg/jwt"
	"shop-smart-api/internal/service/user"
)

type userService struct {
	finder    user.Finder
	collector user.Collector
	modifier  user.Modifier
	creator   user.Creator
	jwtManger jwt.Manager
}

func CreateUserService(
	f user.Finder,
	cs user.Collector,
	m user.Modifier,
	c user.Creator,
	j jwt.Manager,
) UserService {
	return &userService{f, cs, m, c, j}
}

func (uc *userService) Get(id string) (*entity.User, error) {
	return uc.finder.Find(id)
}

func (uc *userService) GetByPhone(phone string) (*entity.User, error) {
	return uc.finder.FindByPhone(phone)
}

func (uc *userService) GetByEmail(email string) (*entity.User, error) {
	return uc.finder.FindByEmail(email)
}

func (uc *userService) GetByOrganization(id int64) ([]*entity.User, error) {
	return uc.finder.FindByOrganization(id)
}

func (uc *userService) ProvideOrCreate(resource string, channel *types.Channel, role string) (*entity.User, string, error) {
	if channel.IsEmail() {
		u, err := uc.finder.FindByEmail(resource)
		if err != nil {
			createdUser, err := uc.creator.Create(resource, channel)
			log.Println(createdUser)
			if err != nil {
				log.Println(err)
				return nil, "", err
			}

			token, err := uc.jwtManger.Generate(createdUser, false, role)
			return u, token, err
		}

		token, err := uc.jwtManger.Generate(u, false, role)
		return u, token, err
	}
	if channel.IsCode() {
		u, err := uc.finder.FindByCode(resource)
		if err != nil {
			createdUser, err := uc.creator.Create(resource, channel)
			log.Println(createdUser)
			if err != nil {
				log.Println(err)
				return nil, "", err
			}

			token, err := uc.jwtManger.Generate(createdUser, true, role)
			return u, token, err
		}

		token, err := uc.jwtManger.Generate(u, true, role)
		return u, token, err
	}
	u, err := uc.finder.FindByPhone(resource)
	if err != nil {
		createdUser, err := uc.creator.Create(resource, channel)
		log.Println(createdUser)
		if err != nil {
			log.Println(err)
			return nil, "", err
		}

		token, err := uc.jwtManger.Generate(createdUser, false, role)
		return createdUser, token, err
	}

	token, err := uc.jwtManger.Generate(u, false, role)
	return u, token, err
}

func (uc *userService) Authenticate(user *entity.User) (string, error) {
	return uc.jwtManger.Generate(user, true, "")
}

func (uc *userService) Update(user *entity.User, email string) (*entity.User, error) {
	return uc.modifier.UpdateUser(user, email)
}
