package entity

type StudentCourse struct{
	Id int `json:"id"`
	StudentID int `json:"studentID"`
	CourseName string `json:"courseName"`
}