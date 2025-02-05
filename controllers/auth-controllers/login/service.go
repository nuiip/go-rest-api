package login

import (
	model "nuiip/go-rest-api/models"
)

type Service interface {
	LoginService(input *InputLogin) (*model.User, string)
}

type service struct {
	repository Repository
}

func NewServiceLogin(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) LoginService(input *InputLogin) (*model.User, string) {

	user := model.User{
		Username: input.Username,
		Password: input.Password,
	}

	resultLogin, errLogin := s.repository.LoginRepository(&user)

	return resultLogin, errLogin
}
