package router

import (
	"database/sql"
	"mySqlAPI/models/student"
	"mySqlAPI/models/subject"
	"mySqlAPI/models/teacher"
	"net/http"
)

// ModelRouter struct
type modelRouter struct {
	db *sql.DB
}

// NewModelRouter constructor
func newModelRouter(db *sql.DB) *modelRouter {
	return &modelRouter{db: db}
}

//TeacherPage function receiver for NewModelRouter
func (mr *modelRouter) teacherPage() func(http.ResponseWriter, *http.Request) {
	return teacher.RouteTeacher(mr.db)
}

//SubjectPage view teacher pagefunction receiver for NewModelRouter
func (mr *modelRouter) subjectPage() func(http.ResponseWriter, *http.Request) {
	return subject.RouteSubject(mr.db)
}

//StudentPage function receiver for NewModelRouter
func (mr *modelRouter) studentPage() func(http.ResponseWriter, *http.Request) {
	return student.RouteStudent(mr.db)
}
