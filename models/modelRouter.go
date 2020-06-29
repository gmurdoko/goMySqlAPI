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

//TeacherPage function receiver for NewModelRouter
func (mr *ModelRouter) TeacherPage() func(http.ResponseWriter, *http.Request) {
	return teacher.RouteTeacher(mr.db)
}

//SubjectPage view teacher pagefunction receiver for NewModelRouter
func (mr *ModelRouter) SubjectPage() func(http.ResponseWriter, *http.Request) {
	return subject.RouteSubject(mr.db)
}

//StudentPage function receiver for NewModelRouter
func (mr *ModelRouter) StudentPage() func(http.ResponseWriter, *http.Request) {
	return student.RouteStudent(mr.db)
}
