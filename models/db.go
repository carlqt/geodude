package models

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

var db *sql.DB
var dbConfig DatabaseConfig

type Application struct {
	Development DevelopmentBody
}

type DevelopmentBody struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Username string
	Password string
	Name     string
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getConfig() string {
	// get the abs path of application.yml
	filepath, _ := filepath.Abs("application.yml")

	// read from the yaml file
	yamlData, err := ioutil.ReadFile(filepath)
	checkErr(err)

	config := new(Application)

	// unmarshal it
	err = yaml.Unmarshal(yamlData, &config)
	checkErr(err)

	dbConfig = config.Development.Database

	return fmt.Sprintf("user=%s dbname=%s sslmode=disable", dbConfig.Username, dbConfig.Name)
}

func InitDB() {
	var err error

	dbInfo := getConfig()

	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
}
