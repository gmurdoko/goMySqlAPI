package student

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

//RouteStudent route student
func RouteStudent(ac *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			studentsRaw := GetAllStudent(ac)
			studentsData, err := json.Marshal(studentsRaw)
			errHandling(err)
			w.Header().Set("content-type", "application/json")
			w.Write([]byte(studentsData))
			fmt.Println("Endpoint hit: GetStudents")
			// case "POST":
			// 	var s domains.Student
			// 	json.NewDecoder(r.Body).Decode(&s)
			// 	fmt.Println("Student")
			// 	fmt.Println(s)
		}
	}
}
