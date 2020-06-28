package student

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"mySqlAPI/utils"
	"net/http"
	"regexp"
)

//RouteStudent route student
func RouteStudent(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getStudents(w, r, db)
		case "POST":
			postStudents(w, r, db)
		case "PUT":
			putStudents(w, r, db)
		case "DELETE":
			delStudents(w, r, db)
		}
	}
}

func getStudents(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var studentsResponse utils.Response
	ID := r.URL.Query()
	if len(ID) > 0 {
		idMap := ID["id"]
		idIsi := idMap[0]
		isIDValid := validasiID(idIsi)
		ide, isIDExist := searchByID(db, idIsi)
		if !isIDValid {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid ID"))
		} else if !isIDExist {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("ID didn't Exist"))
		} else {
			studentData := studentByID(db, ide)
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
	_ = json.NewDecoder(r.Body).Decode(&inStudent)
	updateStudent(db, inStudent.ID, inStudent.FistName, inStudent.LastName, inStudent.Email)
	studentResponse.Status = http.StatusOK
	studentResponse.Message = "Put Success"
	studentResponse.Data = ""
	w.Header().Set("content-type", "application/json")
	byteOfStudent, err := json.Marshal(studentResponse)
	errHandling(err)
	w.Write([]byte(byteOfStudent))
	fmt.Println("Endpoint hit: PutStudent")
}

func postStudents(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var studentsResponse utils.Response
	var inStudent Students
	_ = json.NewDecoder(r.Body).Decode(&inStudent)
	insertStudent(db, inStudent.ID, inStudent.FistName, inStudent.LastName, inStudent.Email)
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
	ide, isIDExist := searchByID(db, ID)
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

func validasiID(ID string) (status bool) {
	regex, _ := regexp.Compile(`[0-9]+`)
	var resID = regex.MatchString(ID)
	if resID != true {
		status = false
	} else {
		status = true
	}
	return status
}

func searchByID(db *sql.DB, id string) (iid string, status bool) {
	err := db.QueryRow("SELECT id FROM students WHERE id = ?;", id).Scan(&iid)
	if err != nil {
		status = false
		return id, status
	}
	status = true
	return iid, status

}
