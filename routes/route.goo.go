package route

import (
	generateGoo "nuiip/go-rest-api/controllers/goo-controllers/goo"
	handlerGoo "nuiip/go-rest-api/handlers/goo-handlers/goo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitGooRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Goo
	*/
	GooRepository := generateGoo.NewRepositoryGoo(db)
	gooService := generateGoo.NewServiceGoo(GooRepository)
	gooHandler := handlerGoo.NewHandlerGoo(gooService)

	/**
	@description All Goo Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/goo", gooHandler.GooHandler)

}
