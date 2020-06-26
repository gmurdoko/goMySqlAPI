package config

import (
	"fmt"
	"mySqlAPI/utils"
)

var (
	//DbEngine,
	dbUser,
	dbPassword,
	dbHost,
	dbPort,
	dbSchema string
)

//EnvConn is Envconnection to db
func EnvConn() (dbEngine, dbSource string) {
	dbEngine = utils.ViperGetEnv("DB_ENGINE", "mysql") //mysql
	dbUser = utils.ViperGetEnv("DB_USER", "root")      //root
	dbPassword = utils.ViperGetEnv("DB_PASSWORD", "toor")
	dbHost = utils.ViperGetEnv("DB_HOST", "localhost") //localhost
	dbPort = utils.ViperGetEnv("DB_PORT", "3306")      //3306
	dbSchema = utils.ViperGetEnv("DB_SCHEMA", "dbenigmaschool")

	dbSource = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbSchema)
	return dbEngine, dbSource
}
