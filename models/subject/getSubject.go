package subject

import (
	"encoding/json"
	"fmt"
	"log"
	"mySqlAPI/config"
	"mySqlAPI/utils"
	"net/http"
)

//Subjects is a struct for json
type Subjects struct {
	SubjectName string `json:"subjectname"`
}

//SubjectsPage is a function to get data of Subject
func SubjectsPage(w http.ResponseWriter, r *http.Request) {
	listSubjects := getAllSubject()
	byteOfSubjects, err := json.Marshal(listSubjects)
	errHandling(err)
	w.Header().Set("Content-Type", "aplication/json")
	w.Write(byteOfSubjects)
	fmt.Println("Endpoint Hit: SubjectsPage")
}

func getAllSubject() []Subjects {
	dbEngine, dbSource := config.EnvConn()
	db, err := utils.ConnDB(dbEngine, dbSource)
	errHandling(err)
	defer db.Close()
	data, err := db.Query("SELECT subject_name FROM subjects;")
	errHandling(err)
	defer data.Close()
	var result = []Subjects{}
	for data.Next() {
		var subject = Subjects{}
		var err = data.Scan(&subject.SubjectName)
		errHandling(err)
		result = append(result, subject)
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
