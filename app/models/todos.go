package models

import (
	"log"
	"time"
)

type Todo struct {
	ID         int
	Title      string
	UserID     int
	CategoryID int
	StartTime  time.Time
	EndTime    time.Time
	CreatedAt  time.Time
}

func (u *User) CreateTodo(title string, category *Category) (err error) {
	cmd := `insert into todos (
		title,
		user_id,
		category_id,
		start_time,
		end_time,
		created_at) values (?, ?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd, title, u.ID, category.ID, time.Now(), time.Now(), time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func GetTodo(id int) (todo Todo, err error) {
	cmd := `select id, title, user_id, category_id, start_time, end_time, created_at from todos where id = ?`
	todo = Todo{}
	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Title,
		&todo.UserID,
		&todo.CategoryID,
		&todo.StartTime,
		&todo.EndTime,
		&todo.CreatedAt)

	return todo, err
}

func GetTodos() (todos []Todo, err error) {
	cmd := `select id, title, user_id, category_id, start_time, end_time, created_at from todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.UserID,
			&todo.CategoryID,
			&todo.StartTime,
			&todo.EndTime,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}

		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}

func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `select id, title, user_id, category_id, start_time, end_time, created_at from todos where user_id = ?`
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.UserID,
			&todo.CategoryID,
			&todo.StartTime,
			&todo.EndTime,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}

		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}

func (t *Todo) UpdateTodo() (err error) {
	cmd := `update todos set title = ?, start_time = ?, end_time = ?`
	_, err = Db.Exec(cmd, t.Title, t.StartTime, t.EndTime)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (t *Todo) DeleteTodo() (err error) {
	cmd := `delete from todos where id = ?`
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
