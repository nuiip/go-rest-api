package model

import (
	"fmt"
	"time"

	util "nuiip/go-rest-api/utils"

	"gorm.io/gorm"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID                 int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username           string         `gorm:"column:username;not null" json:"username"`
	AuthKey            string         `gorm:"column:auth_key;not null" json:"auth_key"`
	Password           string         `gorm:"column:password_hash;not null" json:"password_hash"`
	PasswordResetToken string         `gorm:"column:password_reset_token" json:"password_reset_token"`
	Email              string         `gorm:"column:email;not null" json:"email"`
	Status             int32          `gorm:"column:status;not null;default:1" json:"status"`
	VerificationToken  string         `gorm:"column:verification_token" json:"verification_token"`
	RefreshToken       string         `gorm:"column:refresh_token" json:"refresh_token"`
	Note               string         `gorm:"column:note" json:"note"`
	CreatedAt          time.Time      `gorm:"column:created_at" json:"created_at"`
	CreatedBy          int32          `gorm:"column:created_by" json:"created_by"`
	UpdatedAt          time.Time      `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy          int32          `gorm:"column:updated_by" json:"updated_by"`
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy          int32          `gorm:"column:deleted_by" json:"deleted_by"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	fmt.Println("p")
	fmt.Println(u.Password)
	u.Password = util.HashPassword(u.Password)
	u.CreatedAt = time.Now().Local()
	return nil
}

func (u *User) BeforeUpdate(db *gorm.DB) error {
	u.UpdatedAt = time.Now().Local()
	return nil
}
