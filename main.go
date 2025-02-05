package main

import (
	"log"

	config "nuiip/go-rest-api/configs"
	route "nuiip/go-rest-api/routes"
	util "nuiip/go-rest-api/utils"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	/**
	@description Setup Server
	*/
	router := SetupRouter()
	/**
	@description Run Server
	*/
	log.Fatal(router.Run(":" + util.GodotEnv("GO_PORT")))
}

func SetupRouter() *gin.Engine {
	/**
	@description Setup Database Connection
	*/
	db := config.Connection()
	/**
	@description Init Router
	*/
	router := gin.Default()
	/**
	@description Setup Mode Application
	*/
	if util.GodotEnv("GO_ENV") != "PROD" && util.GodotEnv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if util.GodotEnv("GO_ENV") == "DEV" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	/**
	@description Setup Middleware
	*/
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))
	router.Use(helmet.Default())
	router.Use(gzip.Gzip(gzip.BestCompression))
	/**
	@description Init All Route
	*/
	route.InitGooRoutes(db, router)
	route.InitAuthRoutes(db, router)
	route.InitStudentRoutes(db, router)
	route.InitUserRoutes(db, router)

	return router
}
