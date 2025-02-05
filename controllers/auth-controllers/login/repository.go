package login

import (
	"fmt"
	model "nuiip/go-rest-api/models"
	util "nuiip/go-rest-api/utils"

	"gorm.io/gorm"
)

type Repository interface {
	LoginRepository(input *model.User) (*model.User, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryLogin(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) LoginRepository(input *model.User) (*model.User, string) {

	var users model.User
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	users.Email = input.Email
	users.Password = input.Password

	checkUserAccount := db.Debug().Select("*").Where("username = ?", input.Username).Find(&users)

	if checkUserAccount.RowsAffected < 1 {
		errorCode <- "LOGIN_NOT_FOUND_404"
		return &users, <-errorCode
	}

	// if !users.Status {
	// 	errorCode <- "LOGIN_NOT_ACTIVE_403"
	// 	return &users, <-errorCode
	// }

	fmt.Println(input.Password)
	fmt.Println(users.Password)

	comparePassword := util.ComparePassword(users.Password, input.Password)

	if comparePassword != nil {
		errorCode <- "LOGIN_WRONG_PASSWORD_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
