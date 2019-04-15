package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Connect setup a connection with the database and return the coonection
func Connect() (*sql.DB, error) {

	db, err := sql.Open("postgres", "user=postgres dbname=fatsoma password=SailIn1985! sslmode=disable")

	if err != nil {
		return nil, errors.New("Could not connect to database")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.New("Could not connect to database: " + err.Error())
	}

	return db, nil
}

// App holds the data fo the router and database connection fo testing purpose
type App struct {
	DB     *sql.DB
	Router *mux.Router
}

// Initialize Method to test the conection with the database
func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbname, password)
	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
}
