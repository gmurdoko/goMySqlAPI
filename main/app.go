package main

import (
	"fmt"
	"log"
	"mySqlAPI/config"
	"mySqlAPI/models"
	"net/http"
)

func main() {
	var db = config.EnvConn()
	ac := models.NewAppController(db)
	//Endpoint
	http.HandleFunc("/students", ac.StudentPage())
	http.HandleFunc("/teachers", ac.TeacherPage())
	http.HandleFunc("/subjects", ac.SubjectPage())

	//Server
	fmt.Println("Running On Port 3000")
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
