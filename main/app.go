package main

import (
	"mySqlAPI/config"
)

func main() {
	var db = config.EnvConn()
	appRouter(db)

}
