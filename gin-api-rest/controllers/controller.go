package controllers

import (
	"gin-api-rest/database"
	"gin-api-rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(200, students)
}

func GetStudentById(c *gin.Context) {
	var student models.Student

	id := c.Param("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func GetStudentByCpf(c *gin.Context) {
	var student models.Student

	cpf := c.Param("cpf")
	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	//tries to bind the given json into given struct
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidateData(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&student)
	c.JSON(http.StatusCreated, student)
}

func DeleteStudent(c *gin.Context) {
	var student models.Student

	id := c.Param("id")
	database.DB.Delete(&student, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Aluno deletado com sucesso",
	})
}

func UpdateStudent(c *gin.Context) {
	var input models.Student
	var student models.Student

	id := c.Param("id")

	//validates if it is a valid body
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidateData(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Aluno não encontrado",
		})
		return
	}

	database.DB.Model(&student).Updates(models.Student{
		Name: input.Name,
		CPF:  input.CPF,
		RG:   input.RG,
	})

	c.JSON(http.StatusOK, student)
}

func RenderView(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"students": students,
	})
}

func NotFoundRoute(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
