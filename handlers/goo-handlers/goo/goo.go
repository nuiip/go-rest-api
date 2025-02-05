package handler

import (
	"fmt"
	"net/http"

	Goo "nuiip/go-rest-api/controllers/goo-controllers/goo"
	util "nuiip/go-rest-api/utils"

	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type handler struct {
	service Goo.Service
}

func NewHandlerGoo(service Goo.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GooHandler(ctx *gin.Context) {

	var input Goo.InputGoo
	ctx.ShouldBindJSON(&input)

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Table",
				Message: "table is required on body",
			},
		},
	}

	errResponse, errCount := util.GoValidator(&input, config.Options)
	fmt.Println(errCount)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	resultGoo, errGoo := h.service.GooService(&input)
	util.APIResponse(ctx, errGoo, http.StatusBadRequest, http.MethodPost, resultGoo)
}
