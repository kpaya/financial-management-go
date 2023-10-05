package database

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var once sync.Once

type DbInstance struct {
	Db *sql.DB
}

var dbInstance *DbInstance

func NewDbInstance() *sql.DB {

	if dbInstance == nil {
		once.Do(func() {
			dbInstance = new(DbInstance)
			var err error
			dbInstance.Db, err = sql.Open("sqlite3", "../../../database_test.db")
			if err != nil {
				log.Panicf("create database was not possible\nError: %s", err.Error())
			}

			dbQuery := `
			CREATE TABLE IF NOT EXISTS user (id TEXT PRIMARY KEY, email TEXT, password TEXT, active BOOLEAN);
			CREATE TABLE IF NOT EXISTS transactions (id TEXT PRIMARY KEY, balance REAL, type TEXT);
			CREATE TABLE IF NOT EXISTS wallet (id TEXT PRIMARY KEY, user_id TEXT, active BOOLEAN, FOREIGN KEY(user_id) REFERENCES user(id));
			`

			_, err = dbInstance.Db.Exec(dbQuery)

			if err != nil {
				log.Panicf("creating tables was not possible\nError: %s", err.Error())
			}
		})
	}

	return dbInstance.Db
}
