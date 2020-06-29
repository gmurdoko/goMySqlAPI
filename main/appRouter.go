package main

import (
	"database/sql"
	"fmt"
	"log"
	"mySqlAPI/models"
	"net/http"
)

func appRouter(db *sql.DB) {
	mr := models.NewModelRouter(db)
	//Endpoint
	http.HandleFunc("/students", mr.StudentPage())
	http.HandleFunc("/teachers", mr.TeacherPage())
	http.HandleFunc("/subjects", mr.SubjectPage())
	//Server
	fmt.Println("Running On Port 3000")
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
