package subject

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"mySqlAPI/utils"
	"net/http"
	"strconv"
)

func getSubjects(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var subjectResponse utils.Response
	ID := r.URL.Query()
	if len(ID) > 0 {
		idMap := ID["id"]
		idIsi := idMap[0]
		isIDValid := validasiID(idIsi)
		ide, isIDExist := validateSubjectID(db, idIsi)
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
	var inSubject Subject
	_ = json.NewDecoder(r.Body).Decode(&inSubject)
	insertSubject(db, inSubject.SubjectName)
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
	err := json.NewDecoder(r.Body).Decode(&inSubject)
	ID := strconv.Itoa(inSubject.ID)
	// isIDValid := validasiID(ID)
	ide, isIDExist := validateSubjectID(db, ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID"))
	} else if !isIDExist {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID didn't Exist"))
	} else {
		updateSubject(db, ide, inSubject.SubjectName)
		subjectResponse.Status = http.StatusOK
		subjectResponse.Message = "Put Success"
		subjectResponse.Data = ""
		w.Header().Set("content-type", "application/json")
		byteOfSubject, err := json.Marshal(subjectResponse)
		errHandling(err)
		w.Write([]byte(byteOfSubject))
		fmt.Println("Endpoint hit: PutSubject")
	}
}

func delSubjects(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var subjectResponse utils.Response
	ID := r.FormValue("id")
	isIDValid := validasiID(ID)
	ide, isIDExist := validateSubjectID(db, ID)
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

func validateSubjectID(db *sql.DB, id string) (iid string, status bool) {
	err := db.QueryRow("SELECT id FROM subjects WHERE id = ?;", id).Scan(&iid)
	if err != nil {
		status = false
		return id, status
	}
	status = true
	return iid, status

}
