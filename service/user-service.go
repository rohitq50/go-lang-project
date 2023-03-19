package service

import "github.com/rohitq50/go-lang-project/entity"

type UserService interface {
	Save(entity.User) entity.User
	FindAll() []entity.User
}

type userService struct {
	user []entity.User
}

func New() UserService {
	return &userService{}
}

func (service *userService) Save(user entity.User) entity.User {
	service.user = append(service.user, user)
	return user
}

func (service *userService) FindAll() []entity.User {
	return service.user
}
