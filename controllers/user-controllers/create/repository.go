package createUser

import (
	model "nuiip/go-rest-api/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreateUserRepository(input *model.User) (*model.User, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateUserRepository(input *model.User) (*model.User, string) {

	var users model.User
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	checkUserExist := db.Debug().Select("*").Where("username = ?", input.Username).Find(&users)

	if checkUserExist.RowsAffected > 0 {
		errorCode <- "CREATE_USER_CONFLICT_409"
		return &users, <-errorCode
	}

	users.Username = input.Username
	users.Email = input.Email
	users.Password = input.Password

	addNewUser := db.Debug().Create(&users)
	db.Commit()

	if addNewUser.Error != nil {
		errorCode <- "CREATE_STUDENT_FAILED_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
