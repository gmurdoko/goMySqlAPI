package models

import (
	"database/sql"
	"mySqlAPI/models/student"
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
	return func(w http.ResponseWriter, r *http.Request) {
		// switch r.Method {
		// case "GET":
		// 	subjectsRaw := subject.GetAllSubject(ac.db)
		// 	subjectsData, err := json.Marshal(subjectsRaw)
		// 	errHandling(err)
		// 	w.Header().Set("content-type", "application/json")
		// 	w.Write([]byte(subjectsData))
		// 	fmt.Println("Endpoint hit: GetSubjects")
		// 	case "POST":
		// 		var s domains.Subject
		// 		json.NewDecoder(r.Body).Decode(&s)
		// 		fmt.Println("Subject")
		// 		fmt.Println(s)
		// }
	}
}

//StudentPage view teacher page
func (ac *AppController) StudentPage() func(http.ResponseWriter, *http.Request) {
	return student.RouteStudent(ac.db)
}
