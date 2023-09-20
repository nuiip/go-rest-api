package handlerResultsStudent

import (
	"net/http"

	resultsStudent "nuiip/go-rest-api/controllers/student-controllers/results"
	util "nuiip/go-rest-api/utils"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service resultsStudent.Service
}

func NewHandlerResultsStudent(service resultsStudent.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ResultsStudentHandler(ctx *gin.Context) {

	resultsStudent, errResultsStudent := h.service.ResultsStudentService()

	switch errResultsStudent {

	case "RESULTS_STUDENT_NOT_FOUND_404":
		util.APIResponse(ctx, "Students data is not exists", http.StatusConflict, http.MethodPost, nil)

	default:
		util.APIResponse(ctx, "Results Students data successfully", http.StatusOK, http.MethodPost, resultsStudent)
	}
}
