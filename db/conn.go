package db

import (
	"database/sql"
	"examSearch/configs"
	"fmt"
)

// OpenConnection opens a connection to the database.
func OpenConnection() (*sql.DB, error) {

	// get the database configuration
	conf := configs.GetDB()

	str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disabel", conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	// open the connection
	db, err := sql.Open("postgres", str)
	if err != nil {
		panic(err)
	}

	// check if the connection is alive
	err = db.Ping()

	return db, err
}
