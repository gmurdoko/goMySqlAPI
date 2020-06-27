package student

// //StudentsPageByID is seraching student by id
// func StudentsPageByID(w http.ResponseWriter, r *http.Request, idStudent int) {
// 	listStudents := getStudentByID(idStudent)
// 	byteOfStudents, err := json.Marshal(listStudents)
// 	errHandling(err)
// 	w.Header().Set("Content-Type", "aplication/json")
// 	w.Write(byteOfStudents)
// 	fmt.Println("Endpoint Hit: StudentsPage")
// }

// func getStudentByID(idStudent int) []Students {
// 	dbEngine, dbSource := config.EnvConn()
// 	db, err := utils.ConnDB(dbEngine, dbSource)
// 	errHandling(err)
// 	defer db.Close()
// 	data, err := db.Query("SELECT first_name, last_name, email FROM students where id=?;", idStudent)
// 	errHandling(err)
// 	defer data.Close()
// 	var result = []Students{}
// 	for data.Next() {
// 		var student = Students{}
// 		var err = data.Scan(&student.FistName, &student.LastName, &student.Email)
// 		errHandling(err)
// 		result = append(result, student)
// 	}
// 	if err = data.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 	return result
// }
