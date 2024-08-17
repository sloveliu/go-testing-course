package main

import (
	"log"
	"os"
	"testing"
	"webapp/pkg/db"
)

var app application

// 運行 go test . 時，會先運行
func TestMain(m *testing.M) {
	pathToTemplate = "./../../templates/"
	app.Session = getSession()
	app.DSN = "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=Asia/Taipei connect_timeout=5"
	conn, err := app.connectToDb()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	app.DB = db.PostgresConn{DB: conn}

	os.Exit(m.Run())
}
