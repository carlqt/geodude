package main

import(
  "fmt"
  "database/sql"
  _ "github.com/lib/pq"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "path/filepath"
)

type Application struct {
  Development DevelopmentBody
}

type DevelopmentBody struct {
  Database DatabaseConfig
}

type DatabaseConfig struct {
  Username string
  Password string
  Name string
}

func checkErr(err error) {
  if err != nil {
    panic("Something happened: ", err)
  }
}

var dbConfig DatabaseConfig
var db *sql.DB

func init() {
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
}

func main() {
  var err error
  // open the db
  dbInfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable", dbConfig.Username, dbConfig.Name)
  db, err = sql.Open("postgres", dbInfo)
  checkErr(err)
  defer db.Close()


  // prepare sql statement
  stmnt, err := db.Prepare("INSERT INTO properties(address, latitude, longitude) VALUES($1, $2, $3)")
  checkErr(err)
  defer stmnt.Close()

  //insert data NOW
  _, err = stmnt.Exec("Bugis Junction",1.29960 ,103.85513)
  checkErr(err)
  stmnt.Exec("Bugis Plus",1.30106 ,103.85599)
  stmnt.Exec("Rochor Centre",1.30256 ,103.85487)
  stmnt.Exec("Parkview Square",1.30019 ,103.85761)
  fmt.Println("Data inserted")
}