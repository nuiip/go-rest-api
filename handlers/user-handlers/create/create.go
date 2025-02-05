package handlerCreateStudent

import (
	"net/http"

	createUser "nuiip/go-rest-api/controllers/user-controllers/create"
	util "nuiip/go-rest-api/utils"

	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type handler struct {
	service createUser.Service
}

func NewHandlerCreateUser(service createUser.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateUserHandler(ctx *gin.Context) {

	var input createUser.InputCreateUser
	ctx.ShouldBindJSON(&input)

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Username",
				Message: "username is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Email",
				Message: "email is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "email",
				Field:   "Email",
				Message: "email format is not valid",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "PasswordHash",
				Message: "password is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "gte",
				Field:   "Password",
				Message: "password minimum must be 8 character",
			},
		},
	}

	errResponse, errCount := util.GoValidator(&input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	_, errCreateUser := h.service.CreateUserService(&input)

	switch errCreateUser {

	case "CREATE_STUDENT_CONFLICT_409":
		util.APIResponse(ctx, "Npm student already exist", http.StatusConflict, http.MethodPost, nil)
		return

	case "CREATE_STUDENT_FAILED_403":
		util.APIResponse(ctx, "Create new student account failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		util.APIResponse(ctx, "Create new student account successfully", http.StatusCreated, http.MethodPost, nil)
	}
}
