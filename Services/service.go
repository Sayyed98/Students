package services

import (
	"database/sql"
	"fmt"
	"log"
	entity "student/Entity"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Services interface {
	AddStudent(c *gin.Context, studentName entity.StudentDetails) error
	GetStudent(c *gin.Context, id int) (entity.StudentDetails, error)
}

// const (
// 	dbUser   = "root"
// 	dbName   = "students"
// 	host     = "localhost"
// 	port     = "3306"
// 	password = "12345678"
// 	driver   = "mysql"
// )

type ServiceHandler struct {
	db *sql.DB
}

func NewServiceHandler() (*ServiceHandler, error) {
	// dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, password, host, dbName)
	// dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, password, dbName)
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/students")
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Failed to ping database: %v", err)
		return nil, err
	}

	return &ServiceHandler{db: db}, nil
}

func (s *ServiceHandler) AddStudent(c *gin.Context, studentName entity.StudentDetails) error {
	query := "Insert into studentsDetails(name, phoneNumber) values(?,?)"
	_, err := s.db.Exec(query, studentName.Name, studentName.PhoneNumber)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceHandler) GetStudent(c *gin.Context, id int) (entity.StudentDetails, error) {
	var student entity.StudentDetails
	query := "Select * from studentsDetails where id= ?"
	err := s.db.QueryRow(query, id).Scan(&student.Id, &student.Name, &student.PhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return student, fmt.Errorf("student not found")
		}
		return student, fmt.Errorf("failed to get student: %v", err)
	}
	return student, nil
}
