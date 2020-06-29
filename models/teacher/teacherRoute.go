package teacher

import (
	"database/sql"
	"net/http"
	"regexp"
)

// RouteTeacher routing for model teacher
func RouteTeacher(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getTeachers(w, r, db)
		case "POST":
			postTeachers(w, r, db)
		case "PUT":
			putTeachers(w, r, db)
		case "DELETE":
			delTeachers(w, r, db)
		}
	}
}

func validasiID(ID string) (status bool) {
	regex, _ := regexp.Compile(`[0-9]+`)
	var resID = regex.MatchString(ID)
	if resID != true {
		status = false
	} else {
		status = true
	}
	return status
}
