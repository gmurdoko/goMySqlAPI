package router

import (
	"fmt"
	"log"
	"mySqlAPI/config"
	"mySqlAPI/models"
	"net/http"
)

//AppRouters is router for app.go
func AppRouters() {
	var db = config.EnvConn()
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