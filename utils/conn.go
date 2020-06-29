package utils

import (
	"database/sql"
	"log"

	//package for DB connection using mysql
	_ "github.com/go-sql-driver/mysql"
)

func connCheck(db *sql.DB) (*sql.DB, error) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
	return db, err
}

//ConnDB Connecting app to DB using variabel from EnvConn
func ConnDB(DbEngine, DbSource string) (db *sql.DB, err error) {
	db, _ = sql.Open(DbEngine, DbSource)
	db, err = connCheck(db)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
