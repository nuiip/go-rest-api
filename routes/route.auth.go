package route

import (
	activationAuth "nuiip/go-rest-api/controllers/auth-controllers/activation"
	forgotAuth "nuiip/go-rest-api/controllers/auth-controllers/forgot"
	loginAuth "nuiip/go-rest-api/controllers/auth-controllers/login"
	registerAuth "nuiip/go-rest-api/controllers/auth-controllers/register"
	resendAuth "nuiip/go-rest-api/controllers/auth-controllers/resend"
	resetAuth "nuiip/go-rest-api/controllers/auth-controllers/reset"
	handlerActivation "nuiip/go-rest-api/handlers/auth-handlers/activation"
	handlerForgot "nuiip/go-rest-api/handlers/auth-handlers/forgot"
	handlerLogin "nuiip/go-rest-api/handlers/auth-handlers/login"
	handlerRegister "nuiip/go-rest-api/handlers/auth-handlers/register"
	handlerResend "nuiip/go-rest-api/handlers/auth-handlers/resend"
	handlerReset "nuiip/go-rest-api/handlers/auth-handlers/reset"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Auth
	*/
	LoginRepository := loginAuth.NewRepositoryLogin(db)
	loginService := loginAuth.NewServiceLogin(LoginRepository)
	loginHandler := handlerLogin.NewHandlerLogin(loginService)

	registerRepository := registerAuth.NewRepositoryRegister(db)
	registerService := registerAuth.NewServiceRegister(registerRepository)
	registerHandler := handlerRegister.NewHandlerRegister(registerService)

	activationRepository := activationAuth.NewRepositoryActivation(db)
	activationService := activationAuth.NewServiceActivation(activationRepository)
	activationHandler := handlerActivation.NewHandlerActivation(activationService)

	resendRepository := resendAuth.NewRepositoryResend(db)
	resendService := resendAuth.NewServiceResend(resendRepository)
	resendHandler := handlerResend.NewHandlerResend(resendService)

	forgotRepository := forgotAuth.NewRepositoryForgot(db)
	forgotService := forgotAuth.NewServiceForgot(forgotRepository)
	forgotHandler := handlerForgot.NewHandlerForgot(forgotService)

	resetRepository := resetAuth.NewRepositoryReset(db)
	resetService := resetAuth.NewServiceReset(resetRepository)
	resetHandler := handlerReset.NewHandlerReset(resetService)

	/**
	@description All Auth Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/register", registerHandler.RegisterHandler)
	groupRoute.POST("/login", loginHandler.LoginHandler)
	groupRoute.POST("/activation/:token", activationHandler.ActivationHandler)
	groupRoute.POST("/resend-token", resendHandler.ResendHandler)
	groupRoute.POST("/forgot-password", forgotHandler.ForgotHandler)
	groupRoute.POST("/change-password/:token", resetHandler.ResetHandler)

}
