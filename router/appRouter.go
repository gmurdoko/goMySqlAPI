package router

import (
	"fmt"
	"log"
	"mySqlAPI/config"
	"net/http"
)

//AppRouter is function to route app.go
func AppRouter() {
	var db = config.EnvConn()
	mr := newModelRouter(db)
	//Endpoint
	http.HandleFunc("/students", mr.studentPage())
	http.HandleFunc("/teachers", mr.teacherPage())
	http.HandleFunc("/subjects", mr.subjectPage())
	//Server
	fmt.Println("Running On Port 3000")
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
