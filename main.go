package main

import (
	"log"
	services "student/Services"

	"github.com/gin-gonic/gin"
)

func main() {

	service, err := services.NewServiceHandler()
	if err != nil {
		log.Fatalf("Failed to initialize service handler: %v", err)
	}
	web := services.NewWebHandler(service)
	r := gin.Default()
	r.POST("/add", web.AddStudent)
	r.GET("/student/:id", web.GetStudent)
	r.Run(":8080")
}
