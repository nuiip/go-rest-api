package register

import (
	model "nuiip/go-rest-api/models"
)

type Service interface {
	RegisterService(input *InputRegister) (*model.User, string)
}

type service struct {
	repository Repository
}

func NewServiceRegister(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) RegisterService(input *InputRegister) (*model.User, string) {

	users := model.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	resultRegister, errRegister := s.repository.RegisterRepository(&users)

	return resultRegister, errRegister
}
