package model

type EntityGoo struct {
	Table string `gorm:"type:varchar(255);not null"`
}
