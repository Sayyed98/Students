package services

import (
	"net/http"
	entity "student/Entity"

	"github.com/gin-gonic/gin"
)

type WebHandler struct {
	ser Services
}

func NewWebHandler(w Services) *WebHandler {
	return &WebHandler{ser: w}
}

func (w *WebHandler) AddStudent(c *gin.Context) {
	student := entity.StudentDetails{}

	err := c.ShouldBindJSON(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = w.ser.AddStudent(c, student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student added successfully"})
}

func (w *WebHandler) GetStudent(c *gin.Context) {
	details, err := w.ser.GetStudent(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.json
}
