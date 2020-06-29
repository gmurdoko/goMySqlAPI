package teacher

import (
	"database/sql"
	"net/http"
	"regexp"
)

// RouteTeacher routin;g for teachers
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

func searchByID(db *sql.DB, id string) (iid string, status bool) {
	err := db.QueryRow("SELECT id FROM teachers WHERE id = ?;", id).Scan(&iid)
	if err != nil {
		status = false
		return id, status
	}
	status = true
	return iid, status

}
