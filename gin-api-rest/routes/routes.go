package routes

import (
	"gin-api-rest/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/students", controllers.GetStudents)
	r.POST("/students", controllers.CreateStudent)
	r.GET("/students/:id", controllers.GetStudentById)
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCpf)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.GET("/", controllers.RenderView)

	r.NoRoute(controllers.NotFoundRoute)
	r.Run()
}
