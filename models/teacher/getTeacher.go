package teacher

import (
	"encoding/json"
	"fmt"
	"log"
	"mySqlAPI/config"
	"mySqlAPI/utils"
	"net/http"
)

//Teachers is a struct for json
type Teachers struct {
	FistName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
}

//TeachersPage is a function to get data of Teacher
func TeachersPage(w http.ResponseWriter, r *http.Request) {
	listTeachers := getAllTeacher()
	byteOfTeachers, err := json.Marshal(listTeachers)
	errHandling(err)
	w.Header().Set("Content-Type", "aplication/json")
	w.Write(byteOfTeachers)
	fmt.Println("Endpoint Hit: TeachersPage")
}

func getAllTeacher() []Teachers {
	dbEngine, dbSource := config.EnvConn()
	db, err := utils.ConnDB(dbEngine, dbSource)
	errHandling(err)
	defer db.Close()
	data, err := db.Query("SELECT first_name, last_name, email FROM teachers;")
	errHandling(err)
	defer data.Close()
	var result = []Teachers{}
	for data.Next() {
		var teacher = Teachers{}
		var err = data.Scan(&teacher.FistName, &teacher.LastName, &teacher.Email)
		errHandling(err)
		result = append(result, teacher)
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
