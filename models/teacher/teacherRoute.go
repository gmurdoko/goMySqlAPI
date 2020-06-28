package teacher

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"mySqlAPI/utils"
	"net/http"
	"regexp"
)

// RouteTeacher routin;g for teachers
func RouteTeacher(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getTeachers(w, r, db)
		case "POST":
			postTeachers(w, r, db)
		case "PUT":
			putTeachers(w, r, db)
		case "DELETE":
			delTeachers(w, r, db)
		}
	}
}

func getTeachers(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var teachersResponse utils.Response
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
	insertTeacher(db, inTeacher.ID, inTeacher.FistName, inTeacher.LastName, inTeacher.Email)
	teachersResponse.Status = http.StatusOK
	teachersResponse.Message = "Post Success"
	teachersResponse.Data = ""
	w.Header().Set("content-type", "application/json")
	byteOfTeacher, err := json.Marshal(teachersResponse)
	errHandling(err)
	w.Write([]byte(byteOfTeacher))
	fmt.Println("Endpoint hit: PostTeacher")
}

func putTeachers(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var teachersResponse utils.Response
	var inTeacher Teachers
	_ = json.NewDecoder(r.Body).Decode(&inTeacher)
	updateTeacher(db, inTeacher.ID, inTeacher.FistName, inTeacher.LastName, inTeacher.Email)
	teachersResponse.Status = http.StatusOK
	teachersResponse.Message = "Put Success"
	teachersResponse.Data = ""
	w.Header().Set("content-type", "application/json")
	byteOfTeacher, err := json.Marshal(teachersResponse)
	errHandling(err)
	w.Write([]byte(byteOfTeacher))
	fmt.Println("Endpoint hit: PutTeacher")
}

func delTeachers(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var teachersResponse utils.Response
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
	err := db.QueryRow("SELECT id FROM teachers WHERE id = ?;", id).Scan(&iid)
	if err != nil {
		status = false
		return id, status
	}
	status = true
	return iid, status

}
