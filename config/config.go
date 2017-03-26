package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var mongohost string
var mongouser string
var mongopass string
var mongoport int
var dbname string
var listenport int
var debugmode bool

func Read() {

	fmt.Println("Reading Configuration file...")

	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		fmt.Printf("config.json is missing, have you created it yet?\n", e)
		os.Exit(1)
	}

	var configSettings struct {
		Mongohost  string
		Mongouser  string
		Mongopass  string
		Mongoport  int
		Dbname     string
		Listenport int
		Debugmode  bool
	}

	json.Unmarshal(file, &configSettings)

	mongohost = configSettings.Mongohost
	mongouser = configSettings.Mongouser
	mongopass = configSettings.Mongopass
	mongoport = configSettings.Mongoport
	dbname = configSettings.Dbname
	listenport = configSettings.Listenport
	debugmode = configSettings.Debugmode

	fmt.Println("Configuration file loaded")

}

func GetMongoHost() string {
	return mongohost
}

func GetMongoUser() string {
	return mongouser
}

func GetMongoPass() string {
	return mongopass
}

func GetMongoPort() int {
	return mongoport
}

func GetDBName() string {
	return dbname
}

func GetMongoCreds() (host string, user string, pass string, port int, db string) {

	return mongohost, mongouser, mongopass, mongoport, dbname
}

func GetListenPort() int {
	return listenport
}

func GetDebugMode() bool {
	return debugmode
}
