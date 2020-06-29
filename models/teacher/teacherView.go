package teacher

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"mySqlAPI/utils"
	"net/http"
	"strconv"
)

func getTeachers(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var teachersResponse utils.Response
	ID := r.URL.Query()
	if len(ID) > 0 {
		idMap := ID["id"]
		idIsi := idMap[0]
		isIDValid := validasiID(idIsi)
		ide, isIDExist := validateTeacherID(db, idIsi)
		if !isIDValid {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid ID"))
		} else if !isIDExist {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("ID didn't Exist"))
		} else {
			teacherData := teacherByID(db, ide)
			teachersResponse.Status = http.StatusOK
			teachersResponse.Message = "Get ID Success!"
			teachersResponse.Data = teacherData
			newTeacherData, err := json.Marshal(teachersResponse)
			errHandling(err)
			w.Header().Set("content-type", "application/json")
			w.Write([]byte(newTeacherData))
			fmt.Println("Endpoint hit: GetTeachersByID")
		}

	} else {
		teachersData := getAllTeacher(db)
		teachersResponse.Status = http.StatusOK
		teachersResponse.Message = "Get All Success"
		teachersResponse.Data = teachersData
		newTeachersData, err := json.Marshal(teachersResponse)
		errHandling(err)
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(newTeachersData))
		fmt.Println("Endpoint hit: GetTeachers")
	}
}

func postTeachers(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var teachersResponse utils.Response
	var inTeacher Teachers
	_ = json.NewDecoder(r.Body).Decode(&inTeacher)
	if inTeacher.FistName == "" && inTeacher.LastName == "" && inTeacher.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Can't Read Data or Empty"))
	} else {
		insertTeacher(db, &inTeacher)
		teachersResponse.Status = http.StatusOK
		teachersResponse.Message = "Post Success"
		teachersResponse.Data = ""
		w.Header().Set("content-type", "application/json")
		byteOfTeacher, err := json.Marshal(teachersResponse)
		errHandling(err)
		w.Write([]byte(byteOfTeacher))
		fmt.Println("Endpoint hit: PostTeacher")
	}
}

func putTeachers(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var teachersResponse utils.Response
	var inTeacher Teachers
	err := json.NewDecoder(r.Body).Decode(&inTeacher)
	ID := strconv.Itoa(inTeacher.ID)
	// isIDValid := validasiID(ID)
	_, isIDExist := validateTeacherID(db, ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID"))
	} else if !isIDExist {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID didn't Exist"))
	} else {
		updateTeacher(db, &inTeacher)
		teachersResponse.Status = http.StatusOK
		teachersResponse.Message = "Put Success"
		teachersResponse.Data = ""
		w.Header().Set("content-type", "application/json")
		byteOfTeacher, err := json.Marshal(teachersResponse)
		errHandling(err)
		w.Write([]byte(byteOfTeacher))
		fmt.Println("Endpoint hit: PutTeacher")
	}
}

func delTeachers(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var teachersResponse utils.Response
	ID := r.FormValue("id")
	isIDValid := validasiID(ID)
	ide, isIDExist := validateTeacherID(db, ID)
	if !isIDValid {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID"))
	} else if !isIDExist {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID didn't Exist"))
	} else {
		deleteTeacher(db, ide)
		teachersResponse.Status = http.StatusOK
		teachersResponse.Message = "Delete Success"
		teachersResponse.Data = ""
		w.Header().Set("content-type", "application/json")
		byteOfTeacher, err := json.Marshal(teachersResponse)
		errHandling(err)
		w.Write([]byte(byteOfTeacher))
		fmt.Println("Endpoint hit: DeleteTeacher")
	}
}

func validateTeacherID(db *sql.DB, id string) (iid string, status bool) {
	err := db.QueryRow("SELECT id FROM teachers WHERE id = ?;", id).Scan(&iid)
	if err != nil {
		status = false
		return id, status
	}
	status = true
	return iid, status

}
