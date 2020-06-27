package teacher

import (
	"database/sql"
	"log"
)

//Teachers is a struct for json
type Teachers struct {
	ID       int    `json:"id"`
	FistName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
}

//GetAllTeacher get all data teacher
func GetAllTeacher(db *sql.DB) []Teachers {

	data, err := db.Query("SELECT id, first_name, last_name, email FROM teachers;")
	errHandling(err)
	defer data.Close()
	var result = []Teachers{}
	for data.Next() {
		var teacher = Teachers{}
		var err = data.Scan(&teacher.ID, &teacher.FistName, &teacher.LastName, &teacher.Email)
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
