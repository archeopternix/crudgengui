// Scuffolding for echo framework:  https://echo.labstack.com/
package model

import (
	"database/sql"
	"fmt"
)

type DBConfig struct {
	Environments []Environment `yaml:"environments"`
}

type Environment struct {
	Instance string `yaml:"instance"` // production, development, testing
	Database string `yaml:"database"` // postgres,mysql...
	Host     string `yaml:"host"`     // localhost or IP address
	Port     int    `yaml:"port"`     //5432
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

func InitializeDb(e *Environment) (error, *db.Driver) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		e.Host, e.Port, e.User, e.User, e.Dbname)

	db, err := sql.Open(e.Database, psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return err, db
}

func init() {

	err, Database := InitializeDb()

	if err != nil {
		panic(err)
	}
}
