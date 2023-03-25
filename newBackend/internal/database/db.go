package database

import (
	"database/sql"
	"encoding/json"
	"forum_backend/internal/Log"
	"io/ioutil"

	_ "github.com/mattn/go-sqlite3"
)

type ConfigDb struct {
	Driver string
	Path   string
	Name   string
}

func NewConfDb(l *Log.Logger) *ConfigDb {
	var AppConfig ConfigDb
	raw, err := ioutil.ReadFile("configDB.json")
	if err != nil {
		l.Fatal(err.Error())
	}
	json.Unmarshal(raw, &AppConfig)
	return &AppConfig
}

func (c *ConfigDb) InitDB(l *Log.Logger) *sql.DB {
	db, err := sql.Open(c.Driver, c.Name)
	if err != nil {
		l.Fatal(err.Error())
	}
	if err := db.Ping(); err != nil {
		l.Fatal(err.Error())
	}
	return db
}

func (c *ConfigDb) CreateTables(db *sql.DB, l *Log.Logger) {
	file, err := ioutil.ReadFile("./migrations/db.sql")
	if err != nil {
		l.Fatal(err.Error())
	}
	if _, err := db.Exec(string(file)); err != nil {
		l.Fatal(err.Error())
	}
}
