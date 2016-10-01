package main

import (
	"net/http"

	"github.com/SDkie/employee_graphql_sample/db"
	"github.com/SDkie/employee_graphql_sample/gq"
	log "github.com/Sirupsen/logrus"
)

func main() {
	var err error
	log.SetLevel(log.DebugLevel)

	if err = db.Init("root:pass1234@/graphql_sample_project"); err != nil {
		return
	}
	defer db.Close()

	http.HandleFunc("/graphql", gq.GraphQlHandler)
	if err = http.ListenAndServe(":8080", nil); err != nil {
		log.Errorf("Error while starting webserver %s", err)
	}
}
