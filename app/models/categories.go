package models

import (
	"log"
	"time"
)

type Category struct {
	ID        int
	Name      string
	Color     string
	CreatedAt time.Time
}

func (c *Category) CreateCategory() (err error) {
	cmd := `insert into categories (
		name,
		color,
		created_at) values (?, ?, ?)`

	_, err = Db.Exec(cmd, c.Name, c.Color, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func GetCategory(id int) (category Category, err error) {
	category = Category{}
	cmd := `select id, name, color, created_at from categories where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&category.ID,
		&category.Name,
		&category.Color,
		&category.CreatedAt,
	)

	return category, err
}

func (c *Category) UpdateCategory() (err error) {
	cmd := `update categories set name = ?, color = ? where id = ?`
	_, err = Db.Exec(cmd, c.Name, c.Color, c.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (c *Category) DeleteCategory() (err error) {
	cmd := `delete from categories where id = ?`
	_, err = Db.Exec(cmd, c.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
