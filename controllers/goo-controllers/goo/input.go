package goo

type InputGoo struct {
	Table string `json:"table" validate:"required"`
}
