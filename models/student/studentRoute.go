package student

import (
	"database/sql"
	"net/http"
	"regexp"
)

//RouteStudent route student
func RouteStudent(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getStudents(w, r, db)
		case "POST":
			postStudents(w, r, db)
		case "PUT":
			putStudents(w, r, db)
		case "DELETE":
			delStudents(w, r, db)
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
