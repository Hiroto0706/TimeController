package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"timecontroller/config"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser     = "users"
	tableNameTodo     = "todos"
	tableNameCategory = "categories"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`create table if not exists %s (
		id integer primary key autoincrement,
		uuid string not null unique,
		name string,
		email string,
		password string,
		created_at datetime)`, tableNameUser)

	Db.Exec(cmdU)

	cmdT := fmt.Sprintf(`create table if not exists %s (
		id integer primary key autoincrement,
		title string not null,
		category_id integer,
		start_time datetime,
		end_time datetime,
		created_at datetime)`, tableNameTodo)

	Db.Exec(cmdT)

	cmdC := fmt.Sprintf(`create table if not exists %s (
		id integer primary key autoincrement,
		name string default null,
		color string default null,
		created_at datetime)`, tableNameCategory)

	Db.Exec(cmdC)
}

func CreateUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (crypttext string) {
	crypttext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return crypttext
}
