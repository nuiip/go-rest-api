package route

import (
	createUser "nuiip/go-rest-api/controllers/user-controllers/create"
	handlerCreateUser "nuiip/go-rest-api/handlers/user-handlers/create"
	middleware "nuiip/go-rest-api/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler User
	*/
	createUserRepository := createUser.NewRepositoryCreate(db)
	createUserService := createUser.NewServiceCreate(createUserRepository)
	createUserHandler := handlerCreateUser.NewHandlerCreateUser(createUserService)

	/**
	@description All User Route
	*/
	groupRoute := route.Group("/api/v1").Use(middleware.Auth())
	groupRoute.POST("/user", createUserHandler.CreateUserHandler)
}
