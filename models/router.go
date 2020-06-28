package models

import (
	"database/sql"
	"mySqlAPI/models/student"
	"mySqlAPI/models/subject"
	"mySqlAPI/models/teacher"
	"net/http"
)

// AppController struct
type AppController struct {
	db *sql.DB
}

// NewAppController constructor
func NewAppController(db *sql.DB) *AppController {
	return &AppController{db: db}
}

//TeacherPage view teacher page
func (ac *AppController) TeacherPage() func(http.ResponseWriter, *http.Request) {
	return teacher.RouteTeacher(ac.db)
}

//SubjectPage view teacher page
func (ac *AppController) SubjectPage() func(http.ResponseWriter, *http.Request) {
	return subject.RouteSubject(ac.db)
}

//StudentPage view teacher page
func (ac *AppController) StudentPage() func(http.ResponseWriter, *http.Request) {
	return student.RouteStudent(ac.db)
}
