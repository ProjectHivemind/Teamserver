package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	HOST     = "localhost"
	PORT     = "5432"
	USER     = "hivemind"
	PASSWORD = "changeme123"
	DBNAME   = "hivemind"
)

type DatabaseModel struct {
	db *sql.DB
}

func (d *DatabaseModel) Open() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DBNAME)

	var err error
	d.db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = d.db.Ping()
	if err != nil {
		panic(err)
	}
}

func (d *DatabaseModel) Close() {
	d.db.Close()
}
