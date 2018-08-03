package data

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	file, err := ioutil.ReadFile("./../db.json")
	if err != nil {
		log.Fatal("File error:", err)
		os.Exit(1)
	}
	type dbCredentials struct {
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"db_name"`
		Port     string `json:"port"`
	}
	var credentials dbCredentials
	json.Unmarshal(file, &credentials)

	connection := "user=" + credentials.User +
		" password='" + credentials.Password +
		"' dbname= '" + credentials.DBName +
		"' port= '" + credentials.Port +
		"' sslmode=disable"

	db, err = sql.Open("postgres", connection)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
