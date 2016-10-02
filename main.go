package main

import (
	"net/http"

	"github.com/SDkie/employee_graphql_sample/data"
	"github.com/SDkie/employee_graphql_sample/db"
	"github.com/SDkie/employee_graphql_sample/gq"
	p "github.com/SDkie/employee_graphql_sample/preferences"
	log "github.com/Sirupsen/logrus"
)

func main() {
	var err error
	log.SetLevel(log.DebugLevel)

	p.Init("config.ini")
	if err = db.Init(p.GetMysqlURL()); err != nil {
		return
	}
	defer db.Close()

	data.Init()

	http.HandleFunc("/graphql", gq.GraphQlHandler)
	if err = http.ListenAndServe(p.GetPort(), nil); err != nil {
		log.Errorf("Error while starting webserver %s", err)
	}
}
