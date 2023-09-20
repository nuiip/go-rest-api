package route

import (
	createStudent "nuiip/go-rest-api/controllers/student-controllers/create"
	deleteStudent "nuiip/go-rest-api/controllers/student-controllers/delete"
	resultStudent "nuiip/go-rest-api/controllers/student-controllers/result"
	resultsStudent "nuiip/go-rest-api/controllers/student-controllers/results"
	updateStudent "nuiip/go-rest-api/controllers/student-controllers/update"
	handlerCreateStudent "nuiip/go-rest-api/handlers/student-handlers/create"
	handlerDeleteStudent "nuiip/go-rest-api/handlers/student-handlers/delete"
	handlerResultStudent "nuiip/go-rest-api/handlers/student-handlers/result"
	handlerResultsStudent "nuiip/go-rest-api/handlers/student-handlers/results"
	handlerUpdateStudent "nuiip/go-rest-api/handlers/student-handlers/update"
	middleware "nuiip/go-rest-api/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitStudentRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Student
	*/
	createStudentRepository := createStudent.NewRepositoryCreate(db)
	createStudentService := createStudent.NewServiceCreate(createStudentRepository)
	createStudentHandler := handlerCreateStudent.NewHandlerCreateStudent(createStudentService)

	resultsStudentRepository := resultsStudent.NewRepositoryResults(db)
	resultsStudentService := resultsStudent.NewServiceResults(resultsStudentRepository)
	resultsStudentHandler := handlerResultsStudent.NewHandlerResultsStudent(resultsStudentService)

	resultStudentRepository := resultStudent.NewRepositoryResult(db)
	resultStudentService := resultStudent.NewServiceResult(resultStudentRepository)
	resultStudentHandler := handlerResultStudent.NewHandlerResultStudent(resultStudentService)

	deleteStudentRepository := deleteStudent.NewRepositoryDelete(db)
	deleteStudentService := deleteStudent.NewServiceDelete(deleteStudentRepository)
	deleteStudentHandler := handlerDeleteStudent.NewHandlerDeleteStudent(deleteStudentService)

	updateStudentRepository := updateStudent.NewRepositoryUpdate(db)
	updateStudentService := updateStudent.NewServiceUpdate(updateStudentRepository)
	updateStudentHandler := handlerUpdateStudent.NewHandlerUpdateStudent(updateStudentService)

	/**
	@description All Student Route
	*/
	groupRoute := route.Group("/api/v1").Use(middleware.Auth())
	groupRoute.POST("/student", createStudentHandler.CreateStudentHandler)
	groupRoute.GET("/student", resultsStudentHandler.ResultsStudentHandler)
	groupRoute.GET("/student/:id", resultStudentHandler.ResultStudentHandler)
	groupRoute.DELETE("/student/:id", deleteStudentHandler.DeleteStudentHandler)
	groupRoute.PUT("/student/:id", updateStudentHandler.UpdateStudentHandler)
}
