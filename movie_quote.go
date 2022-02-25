package main

import (
	"database/sql"
	"encoding/json"
	mysql "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type Quote struct {
	Quote string
	Movie string
	Year  int
}

func getQuote(w http.ResponseWriter, req *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(w).Header().Set("Content-Type", "application/json")

	response, _ := json.Marshal(returnCollection())

	(w).Write(response)
}

var db *sql.DB
var cfg mysql.Config

func main() {
	cfg = mysql.Config{
		User:                 "admin",
		Passwd:               "jupiter4",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "movie_quotes",
		AllowNativePasswords: true,
	}

	http.HandleFunc("/movie_quote", getQuote)
	http.ListenAndServe(":8090", nil)
}

func returnCollection() Quote {

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	var (
		quote string
		movie string
		year  int
	)

	err = db.QueryRow("select `quote`, `movie`, `year` from `moviequotes` ORDER BY RAND() LIMIT 1").Scan(&quote, &movie, &year)
	if err != nil {
		log.Fatal(err)
	}

	return Quote{Quote: quote, Movie: movie, Year: year}
}
