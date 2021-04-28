package crud

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// Default variables for connecting to the database
var HOST string = "localhost"
var PORT string = "5432"
var USER string = "hivemind"
var PASSWORD string = "changeme123"
var DBNAME string = "hivemind"
var SSLMODE string = "disabled"

const TimeStamp = time.Stamp

// DatabaseModel holds the database connection object
type DatabaseModel struct {
	db *sql.DB
}

// SetDatabaseOptions is a setter for allowing the database options to be changed.
func SetDatabaseOptions(host, port, user, password, sslMode string) {
	HOST = host
	PORT = port
	USER = user
	PASSWORD = password
	SSLMODE = sslMode
}

// Open creates the database connection
func (d *DatabaseModel) Open() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		HOST, PORT, USER, PASSWORD, DBNAME, SSLMODE)

	var err error
	d.db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = d.db.Ping()
	if err != nil {
		return
	}
}

// Close ends the database connection
func (d *DatabaseModel) Close() {
	d.db.Close()
}
