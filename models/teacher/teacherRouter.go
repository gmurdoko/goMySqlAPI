package teacher

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"mySqlAPI/utils"
	"net/http"
)

// RouteTeacher routing for teachers
func RouteTeacher(ac *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var teacherResponse utils.Response
		switch r.Method {
		case "GET":
			teachersData := GetAllTeacher(ac)
			newTeachersData, err := json.Marshal(teachersData)
			errHandling(err)
			teacherResponse.Message = http.StatusOK
			teacherResponse.Data = teachersData
			w.WriteHeader()
			w.Header().Set("content-type", "application/json")
			w.Write([]byte(teacherResponse))
			fmt.Println("Endpoint hit: GetTeachers")
			// case "POST":
			// 	var t domains.Teacher
			// 	json.NewDecoder(r.Body).Decode(&t)
			// 	fmt.Println("Teacher")
			// 	fmt.Println(t)
		}
	}
}
