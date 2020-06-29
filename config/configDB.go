package config

import (
	"database/sql"
	"fmt"
	"log"
	"mySqlAPI/utils"
)

var (
	dbUser,
	dbPassword,
	dbHost,
	dbPort,
	dbSchema string
)

//EnvConn is function to get Environment Variabel for DB connection
func EnvConn() *sql.DB {
	dbEngine := utils.ViperGetEnv("DB_ENGINE", "mysql") //mysql
	dbUser = utils.ViperGetEnv("DB_USER", "root")       //root
	dbPassword = utils.ViperGetEnv("DB_PASSWORD", "password")
	dbHost = utils.ViperGetEnv("DB_HOST", "localhost") //localhost
	dbPort = utils.ViperGetEnv("DB_PORT", "3306")      //3306
	dbSchema = utils.ViperGetEnv("DB_SCHEMA", "schema")

	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbSchema)
	db, err := utils.ConnDB(dbEngine, dbSource)
	if err != nil {
		log.Panic(err)
	}
	return db
	// return dbEngine, dbSource
}
