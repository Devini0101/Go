package main

import (
	"bytes"
	"encoding/json"
	"gin-api-rest/controllers"
	"gin-api-rest/database"
	"gin-api-rest/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var id int

func SetupRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateStudentMock() {
	student := models.Student{
		Name: "Test Student",
		CPF:  "12345678911",
		RG:   "123456789",
	}

	database.DB.Create(&student)
	id = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, id)
}

func TestListAllStudents(t *testing.T) {
	database.Connect()
	CreateStudentMock()
	r := SetupRoutes()
	defer DeleteStudentMock() //when func ends
	r.GET("/students", controllers.GetStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestSearchStudentByCPF(t *testing.T) {
	database.Connect()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutes()
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCpf)
	req, _ := http.NewRequest("GET", "/students/cpf/12345678911", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetStudentById(t *testing.T) {
	database.Connect()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutes()
	r.GET("/students/:id", controllers.GetStudentById)
	path := "/students/" + strconv.Itoa(id)
	req, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)
	var studentMock models.Student
	json.Unmarshal(response.Body.Bytes(), &studentMock)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Test Student", studentMock.Name, "Os nomes devem ser iguais")
	assert.Equal(t, "12345678911", studentMock.CPF, "O CPF deve ser o mesmo")
	assert.Equal(t, "123456789", studentMock.RG, "O RG deve ser o mesmo")
}

func TestDeleteStudent(t *testing.T) {
	database.Connect()
	CreateStudentMock()
	r := SetupRoutes()
	r.DELETE("/students/:id", controllers.DeleteStudent)

	path := "/students/" + strconv.Itoa(id)
	req, _ := http.NewRequest("DELETE", path, nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestUpdateStudent(t *testing.T) {
	database.Connect()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutes()
	r.PATCH("/students/:id", controllers.UpdateStudent)

	studentMock := models.Student{
		Name: "Test Student",
		CPF:  "47345678911",
		RG:   "123456700",
	}

	jsonStudent, _ := json.Marshal(studentMock)

	path := "/students/" + strconv.Itoa(id)
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(jsonStudent))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	var updatedStudent models.Student
	json.Unmarshal(response.Body.Bytes(), &updatedStudent)

	assert.Equal(t, "Test Student", updatedStudent.Name)
	assert.Equal(t, "47345678911", updatedStudent.CPF)
	assert.Equal(t, "123456700", updatedStudent.RG)
}
