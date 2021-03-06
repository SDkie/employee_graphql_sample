package preferences

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/go-ini/ini"
)

var (
	env      = "dev"
	port     = "8080"
	mysqlUrl = "root:pass1234@/employee_graphql_sample"
)

func GetEnv() string {
	return env
}

func GetPort() string {
	return ":" + port
}

func GetMysqlURL() string {
	return mysqlUrl
}

func Init(iniFile string) {
	var err error
	cfg, err := ini.Load(iniFile)
	if err != nil {
		log.Errorf("Preference Init: Error while loding config.ini file, %s.\nUsing Default values\n", err)
		return
	}

	temp := os.Getenv("ENV")
	if temp != "" {
		env = temp
	}

	config, err := cfg.GetSection(env)
	if err != nil {
		panic(err.Error())
	}

	port = config.Key("PORT").String()
	mysqlUrl = config.Key("MYSQL_URL").String()

	log.Infoln("Preference Init: DONE")
}
