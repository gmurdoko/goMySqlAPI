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

func getAllTeacher(db *sql.DB) []Teachers {
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

func teacherByID(db *sql.DB, id string) Teachers {
	var teacher = Teachers{}
	err := db.QueryRow("SELECT id, first_name, last_name, email FROM teachers WHERE id = ?;", id).Scan(&teacher.ID, &teacher.FistName, &teacher.LastName, &teacher.Email)
	errHandling(err)
	return teacher
}

func insertTeacher(db *sql.DB, inTeacher *Teachers) {
	tx, err := db.Begin()
	errHandling(err)
	// fmt.Sprintf("%v", id)
	_, err = tx.Exec("INSERT INTO teachers(first_name, last_name, email) VALUES(?, ?, ?)", inTeacher.FistName, inTeacher.LastName, inTeacher.Email)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	errHandling(tx.Commit())
}

func updateTeacher(db *sql.DB, inTeacher *Teachers) {
	tx, err := db.Begin()
	errHandling(err)
	_, err = tx.Exec("UPDATE teachers set first_name = ?, last_name = ?, email = ? WHERE id = ?", inTeacher.FistName, inTeacher.LastName, inTeacher.Email, inTeacher.ID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	errHandling(tx.Commit())
}
func deleteTeacher(db *sql.DB, id string) {
	tx, err := db.Begin()
	errHandling(err)

	_, err = tx.Exec("DELETE FROM teachers WHERE id=?;", id)
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
