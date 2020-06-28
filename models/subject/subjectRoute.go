package subject

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"mySqlAPI/utils"
	"net/http"
	"regexp"
)

// RouteSubject routin;g for Subjects
func RouteSubject(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getSubjects(w, r, db)
		case "POST":
			postSubjects(w, r, db)
		case "PUT":
			putSubjects(w, r, db)
		case "DELETE":
			delSubjects(w, r, db)
		}
	}
}

func getSubjects(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var subjectResponse utils.Response
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
			subjectData := subjectByID(db, ide)
			subjectResponse.Status = http.StatusOK
			subjectResponse.Message = "Get ID Success!"
			subjectResponse.Data = subjectData
			newSubjectData, err := json.Marshal(subjectResponse)
			errHandling(err)
			w.Header().Set("content-type", "application/json")
			w.Write([]byte(newSubjectData))
			fmt.Println("Endpoint hit: GetSubjectsByID")
		}

	} else {
		subjectData := getAllSubject(db)
		subjectResponse.Status = http.StatusOK
		subjectResponse.Message = "Get All Success"
		subjectResponse.Data = subjectData
		newSubjectData, err := json.Marshal(subjectResponse)
		errHandling(err)
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(newSubjectData))
		fmt.Println("Endpoint hit: GetSubjects")
	}
}

func postSubjects(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var subjectResponse utils.Response
	var inSubject Subjects
	_ = json.NewDecoder(r.Body).Decode(&inSubject)
	insertSubject(db, inSubject.ID, inSubject.SubjectName)
	subjectResponse.Status = http.StatusOK
	subjectResponse.Message = "Post Success"
	subjectResponse.Data = ""
	w.Header().Set("content-type", "application/json")
	byteOfSubject, err := json.Marshal(subjectResponse)
	errHandling(err)
	w.Write([]byte(byteOfSubject))
	fmt.Println("Endpoint hit: PostSubject")
}

func putSubjects(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var subjectResponse utils.Response
	var inSubject Subjects
	_ = json.NewDecoder(r.Body).Decode(&inSubject)
	updateSubject(db, inSubject.ID, inSubject.SubjectName)
	subjectResponse.Status = http.StatusOK
	subjectResponse.Message = "Put Success"
	subjectResponse.Data = ""
	w.Header().Set("content-type", "application/json")
	byteOfSubject, err := json.Marshal(subjectResponse)
	errHandling(err)
	w.Write([]byte(byteOfSubject))
	fmt.Println("Endpoint hit: PutSubject")
}

func delSubjects(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var subjectResponse utils.Response
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
		deleteSubject(db, ide)
		subjectResponse.Status = http.StatusOK
		subjectResponse.Message = "Delete Success"
		subjectResponse.Data = ""
		w.Header().Set("content-type", "application/json")
		byteOfSubject, err := json.Marshal(subjectResponse)
		errHandling(err)
		w.Write([]byte(byteOfSubject))
		fmt.Println("Endpoint hit: DeleteSubject")
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
	err := db.QueryRow("SELECT id FROM subjects WHERE id = ?;", id).Scan(&iid)
	if err != nil {
		status = false
		return id, status
	}
	status = true
	return iid, status

}
