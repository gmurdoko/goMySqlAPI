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

func getAllStudent(db *sql.DB) []Students {
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

func getStudentByID(db *sql.DB, id string) Students {
	var student = Students{}
	err := db.QueryRow("SELECT id, first_name, last_name, email FROM students WHERE id = ?;", id).Scan(&student.ID, &student.FistName, &student.LastName, &student.Email)
	errHandling(err)
	return student
}

func insertStudent(db *sql.DB, inStudent *Students) {
	tx, err := db.Begin()
	errHandling(err)
	// fmt.Sprintf("%v", id)
	_, err = tx.Exec("INSERT INTO students(first_name, last_name, email) VALUES(?, ?, ?)", inStudent.FistName, inStudent.LastName, inStudent.Email)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	errHandling(tx.Commit())
}

func updateStudent(db *sql.DB, inStudent *Students) {
	tx, err := db.Begin()
	errHandling(err)
	_, err = tx.Exec("UPDATE students set first_name = ?, last_name = ?, email = ? WHERE id = ?", inStudent.FistName, inStudent.LastName, inStudent.Email, inStudent.ID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	errHandling(tx.Commit())
}
func deleteStudent(db *sql.DB, id string) {
	tx, err := db.Begin()
	errHandling(err)

	_, err = tx.Exec("DELETE FROM students WHERE id=?;", id)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	errHandling(tx.Commit())
}
func errHandling(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
