package student

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"mySqlAPI/utils"
	"net/http"
	"strconv"
)

func getStudents(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var studentsResponse utils.Response
	ID := r.URL.Query()
	if len(ID) > 0 {
		idMap := ID["id"]
		idIsi := idMap[0]
		isIDValid := validasiID(idIsi)
		ide, isIDExist := validateStudentID(db, idIsi)
		if !isIDValid {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid ID"))
		} else if !isIDExist {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("ID didn't Exist"))
		} else {
			studentData := getStudentByID(db, ide)
			studentsResponse.Status = http.StatusOK
			studentsResponse.Message = "Get ID Success!"
			studentsResponse.Data = studentData
			newStudentData, err := json.Marshal(studentsResponse)
			errHandling(err)
			w.Header().Set("content-type", "application/json")
			w.Write([]byte(newStudentData))
			fmt.Println("Endpoint hit: GetStudentByID")
		}

	} else {
		studentsData := getAllStudent(db)
		studentsResponse.Status = http.StatusOK
		studentsResponse.Message = "Get All Success"
		studentsResponse.Data = studentsData
		newStudentsData, err := json.Marshal(studentsResponse)
		errHandling(err)
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(newStudentsData))
		fmt.Println("Endpoint hit: GetStudents")
	}
}

func putStudents(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var studentResponse utils.Response
	var inStudent Students
	err := json.NewDecoder(r.Body).Decode(&inStudent)
	ID := strconv.Itoa(inStudent.ID)
	// isIDValid := validasiID(ID)
	ide, isIDExist := validateStudentID(db, ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID"))
	} else if !isIDExist {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID didn't Exist"))
	} else {
		updateStudent(db, ide, inStudent.FistName, inStudent.LastName, inStudent.Email)
		studentResponse.Status = http.StatusOK
		studentResponse.Message = "Put Success"
		studentResponse.Data = ""
		w.Header().Set("content-type", "application/json")
		byteOfStudent, err := json.Marshal(studentResponse)
		errHandling(err)
		w.Write([]byte(byteOfStudent))
		fmt.Println("Endpoint hit: PutStudent")
	}
}

func postStudents(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var studentsResponse utils.Response
	var inStudent Student
	_ = json.NewDecoder(r.Body).Decode(&inStudent)
	insertStudent(db, inStudent.FistName, inStudent.LastName, inStudent.Email)
	studentsResponse.Status = http.StatusOK
	studentsResponse.Message = "Post Success"
	studentsResponse.Data = ""
	w.Header().Set("content-type", "application/json")
	byteOfStudent, err := json.Marshal(studentsResponse)
	errHandling(err)
	w.Write([]byte(byteOfStudent))
	fmt.Println("Endpoint hit: PostStudent")
}

func delStudents(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var studentResponse utils.Response
	ID := r.FormValue("id")
	isIDValid := validasiID(ID)
	ide, isIDExist := validateStudentID(db, ID)
	if !isIDValid {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID"))
	} else if !isIDExist {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID didn't Exist"))
	} else {
		deleteStudent(db, ide)
		studentResponse.Status = http.StatusOK
		studentResponse.Message = "Delete Success"
		studentResponse.Data = ""
		w.Header().Set("content-type", "application/json")
		byteOfStudent, err := json.Marshal(studentResponse)
		errHandling(err)
		w.Write([]byte(byteOfStudent))
		fmt.Println("Endpoint hit: DeleteStudent")
	}
}

func validateStudentID(db *sql.DB, id string) (iid string, status bool) {
	err := db.QueryRow("SELECT id FROM students WHERE id = ?;", id).Scan(&iid)
	if err != nil {
		status = false
		return id, status
	}
	status = true
	return iid, status

}
