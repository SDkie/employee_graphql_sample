package main

import (
	"github.com/SDkie/employee_graphql_sample/db"
	log "github.com/Sirupsen/logrus"
)

func main() {
	var err error
	log.SetLevel(log.DebugLevel)

	if err = db.Init("root:pass1234@/graphql_sample_project"); err != nil {
		return
	}
	defer db.Close()
}
