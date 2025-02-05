package config

import (
	"log"
	"os"

	util "nuiip/go-rest-api/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	databaseURI := make(chan string, 1)

	if os.Getenv("GO_ENV") != "production" {
		databaseURI <- util.GodotEnv("DATABASE_URI_DEV")
	} else {
		databaseURI <- os.Getenv("DATABASE_URI_PROD")
	}

	db, err := gorm.Open(mysql.Open(<-databaseURI), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if os.Getenv("GO_ENV") != "production" {
		logrus.Info("Connection to Database Successfully")
	}

	// err = db.AutoMigrate(
	// 	&model.EntityUsers{},
	// 	&model.EntityStudent{},
	// )

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}
