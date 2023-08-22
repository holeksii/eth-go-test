package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func WaitDB() {
	start := time.Now().Add(time.Second * -10)

	for {
		err := DB.Ping()
		if err == nil {
			break
		}
		if time.Now().After(start.Add(time.Second * 10)) {
			log.Println(err)
			log.Fatal("Waiting for DB connection timeout")
			start = time.Now()
		}
	}
}

func InitDB() {
	connStr := "user=postgres password=secret dbname=ETHDB host=pgdb sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to DB")
	}
}

func CloseDB() {
	DB.Close()
}
