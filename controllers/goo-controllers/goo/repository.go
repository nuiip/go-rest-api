package goo

import (
	config "nuiip/go-rest-api/configs"
	model "nuiip/go-rest-api/models"

	"gorm.io/gen"
	"gorm.io/gorm"
)

type Repository interface {
	GooRepository(input *model.EntityGoo) (*model.EntityGoo, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryGoo(db *gorm.DB) *repository {
	return &repository{db: db}
}

func parseErrorChain(err interface{}) string {
	switch e := err.(type) {
	case error:
		return e.Error() // Standard error case
	default:
		return "Unexpected error"
	}
}

func (r *repository) GooRepository(input *model.EntityGoo) (*model.EntityGoo, string) {

	var goo model.EntityGoo
	db := config.Connection()
	errorCode := make(chan string, 1)
	errorCode <- "nil"

	// Initialize generator
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./models", // Output directory for models
		ModelPkgPath: "models",
	})

	g.UseDB(db) // Use the connected database

	// Generate models from all tables
	// g.GenerateAllTable()

	// Generate model for a specific table (replace "users" with your table name)
	g.GenerateModel(input.Table)
	// if err != nil {
	// 	errorCode <- parseErrorChain(err)
	// }
	// Execute the generator
	g.Execute()

	return &goo, <-errorCode
}
