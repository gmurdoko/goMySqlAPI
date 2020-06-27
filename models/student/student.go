package student

import (
	"database/sql"
	"log"
)

//Students is a struct for json
type Students struct {
	ID       int    `json:"id"`
	FistName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
}

//GetAllStudent is get all student
func GetAllStudent(db *sql.DB) []Students {
	data, err := db.Query("SELECT id, first_name, last_name, email FROM students;")
	errHandling(err)
	defer data.Close()
	var result = []Students{}
	for data.Next() {
		var student = Students{}
		var err = data.Scan(&student.ID, &student.FistName, &student.LastName, &student.Email)
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
