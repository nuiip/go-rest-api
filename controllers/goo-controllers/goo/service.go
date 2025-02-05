package goo

import (
	model "nuiip/go-rest-api/models"
)

type Service interface {
	GooService(input *InputGoo) (*model.EntityGoo, string)
}

type service struct {
	repository Repository
}

func NewServiceGoo(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GooService(input *InputGoo) (*model.EntityGoo, string) {

	goo := model.EntityGoo{
		Table: input.Table,
	}

	resultGoo, errGoo := s.repository.GooRepository(&goo)

	return resultGoo, errGoo
}
