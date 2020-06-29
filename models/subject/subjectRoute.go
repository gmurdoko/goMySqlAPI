package subject

import (
	"database/sql"
	"net/http"
	"regexp"
)

// RouteSubject routin;g for Subjects
func RouteSubject(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getSubjects(w, r, db)
		case "POST":
			postSubjects(w, r, db)
		case "PUT":
			putSubjects(w, r, db)
		case "DELETE":
			delSubjects(w, r, db)
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
