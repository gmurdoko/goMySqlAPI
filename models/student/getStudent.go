package student

import (
	"encoding/json"
	"fmt"
	"log"
	"mySqlAPI/config"
	"mySqlAPI/utils"
	"net/http"
)

//Students is a struct for json
type Students struct {
	FistName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
}

//StudentsPage is a function to get data of student
func StudentsPage(w http.ResponseWriter, r *http.Request) {
	listStudents := getAllStudent()
	byteOfStudents, err := json.Marshal(listStudents)
	errHandling(err)
	w.Header().Set("Content-Type", "aplication/json")
	w.Write(byteOfStudents)
	fmt.Println("Endpoint Hit: StudentsPage")
}

func getAllStudent() []Students {
	dbEngine, dbSource := config.EnvConn()
	db, err := utils.ConnDB(dbEngine, dbSource)
	errHandling(err)
	defer db.Close()
	data, err := db.Query("SELECT first_name, last_name, email FROM students;")
	errHandling(err)
	defer data.Close()
	var result = []Students{}
	for data.Next() {
		var student = Students{}
		var err = data.Scan(&student.FistName, &student.LastName, &student.Email)
		errHandling(err)
		result = append(result, student)
	}
	if err = data.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
func errHandling(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
