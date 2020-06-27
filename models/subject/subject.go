package subject

//Subjects is a struct for json
type Subjects struct {
	ID          int    `json:"id"`
	SubjectName string `json:"subjectName"`
}

// //SubjectsPage is a function to get data of Subject
// func SubjectsPage(w http.ResponseWriter, r *http.Request) {
// 	listSubjects := getAllSubject()
// 	var dataResponse response.Response
// 	dataResponse.Message = "Success!"
// 	dataResponse.Data = listSubjects
// 	byteOfSubjects, err := json.Marshal(dataResponse)
// 	errHandling(err)
// 	w.Header().Set("Content-Type", "aplication/json")
// 	// w.WriteHeader()
// 	w.Write(byteOfSubjects)
// 	fmt.Println("Endpoint Hit: SubjectsPage")
// }

// func GetAllSubject() []Subjects {
// 	// dbEngine, dbSource := config.EnvConn()
// 	// db, err := utils.ConnDB(dbEngine, dbSource)
// 	// errHandling(err)
// 	// defer db.Close()
// 	data, err := db.Query("SELECT id, subject_name FROM subjects;")
// 	errHandling(err)
// 	defer data.Close()
// 	var result = []Subjects{}
// 	for data.Next() {
// 		var subject = Subjects{}
// 		var err = data.Scan(&subject.ID, &subject.SubjectName)
// 		errHandling(err)
// 		result = append(result, subject)
// 	}
// 	if err = data.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 	return result
// }
// func errHandling(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
