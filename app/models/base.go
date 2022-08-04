package models

import (
	"database/sql"
	"fmt"
	"log"
	"timecontroller/config"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	Db, err := sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`create table if not exists %s (
		id integer primary key autoincrement,
		uuid string not null unique,
		name string,
		email string,
		password string,
		start_time datetime,
		end_time datetime,
		created_at datetime)`, tableNameUser)
}
