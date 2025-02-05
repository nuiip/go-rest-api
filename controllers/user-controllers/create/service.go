package createUser

import (
	"fmt"
	model "nuiip/go-rest-api/models"
)

type Service interface {
	CreateUserService(input *InputCreateUser) (*model.User, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateUserService(input *InputCreateUser) (*model.User, string) {

	fmt.Println(input.Password)
	Users := model.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	resultCreateUser, errCreateUser := s.repository.CreateUserRepository(&Users)

	return resultCreateUser, errCreateUser
}
