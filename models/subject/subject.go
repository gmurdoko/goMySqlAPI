package subject

import (
	"database/sql"
	"fmt"
	"log"
)

//Subjects is a struct for json
type Subjects struct {
	ID          int    `json:"id"`
	SubjectName string `json:"subjectName"`
}

func getAllSubject(db *sql.DB) []Subjects {
	data, err := db.Query("SELECT id, subject_name FROM subjects;")
	errHandling(err)
	defer data.Close()
	var result = []Subjects{}
	for data.Next() {
		var subject = Subjects{}
		var err = data.Scan(&subject.ID, &subject.SubjectName)
		errHandling(err)
		result = append(result, subject)
	}
	if err = data.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func subjectByID(db *sql.DB, id string) Subjects {
	var subject = Subjects{}
	err := db.QueryRow("SELECT id, subject_name FROM subjects WHERE id = ?;", id).Scan(&subject.ID, &subject.SubjectName)
	errHandling(err)
	return subject
}

func insertSubject(db *sql.DB, id int, subjectName string) {
	tx, err := db.Begin()
	errHandling(err)
	fmt.Sprintf("%v", id)
	_, err = tx.Exec("INSERT INTO subjects(subject_name) VALUES(?)", subjectName)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	errHandling(tx.Commit())
}

func updateSubject(db *sql.DB, id int, subjectName string) {
	tx, err := db.Begin()
	errHandling(err)
	_, err = tx.Exec("UPDATE subjects set subject_name = ? WHERE id = ?", subjectName, id)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	errHandling(tx.Commit())
}

func deleteSubject(db *sql.DB, id string) {
	tx, err := db.Begin()
	errHandling(err)

	_, err = tx.Exec("DELETE FROM subjects WHERE id=?;", id)
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
