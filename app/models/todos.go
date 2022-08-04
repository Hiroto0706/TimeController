package models

import (
	"fmt"
	"log"
	"time"
)

type Todos struct {
	ID         int
	Title      string
	CategoryID int
	StartTime  time.Time
	EndTime    time.Time
	CreatedAt  time.Time
}

func (u *User) CreateTodo(title string, category *Category) (err error) {
	cmd := `insert into todos (
		title,
		category_id,
		start_time,
		end_time,
		created_at) values (?, ?, ?, ?, ?)`

	fmt.Println(title)
	fmt.Println(category.ID)

	_, err = Db.Exec(cmd, title, category.ID, time.Now(), time.Now(), time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	return err
}
