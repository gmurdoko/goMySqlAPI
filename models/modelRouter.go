package models

import (
	"database/sql"
	"mySqlAPI/models/student"
	"mySqlAPI/models/subject"
	"mySqlAPI/models/teacher"
	"net/http"
)

// ModelRouter struct
type ModelRouter struct {
	db *sql.DB
}

// NewModelRouter constructor
func NewModelRouter(db *sql.DB) *ModelRouter {
	return &ModelRouter{db: db}
}

//TeacherPage view teacher page
func (mr *ModelRouter) TeacherPage() func(http.ResponseWriter, *http.Request) {
	return teacher.RouteTeacher(mr.db)
}

//SubjectPage view teacher page
func (mr *ModelRouter) SubjectPage() func(http.ResponseWriter, *http.Request) {
	return subject.RouteSubject(mr.db)
}

//StudentPage view teacher page
func (mr *ModelRouter) StudentPage() func(http.ResponseWriter, *http.Request) {
	return student.RouteStudent(mr.db)
}
