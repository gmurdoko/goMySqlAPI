package main

import (
	"fmt"
	"log"
	"mySqlAPI/models/student"
	"mySqlAPI/models/subject"
	"mySqlAPI/models/teacher"
	"net/http"
)

func main() {
	//Endpoint
	http.HandleFunc("/students", student.StudentsPage)
	http.HandleFunc("/teachers", teacher.TeachersPage)
	http.HandleFunc("/subjects", subject.SubjectsPage)

	//Server
	fmt.Println("Running On Port 3000")
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
