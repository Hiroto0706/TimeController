package main

import (
	"fmt"
	"timecontroller/app/models"
)

func main() {
	// fmt.Println(config.Config.SQLDriver, config.Config.Port, config.Config.DbName, config.Config.LogFile)

	// log.Println("this is a test")

	fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "ブルーのブチャラティ"
	u.Email = "test@test.com"

	fmt.Println(u)
	u.CreateUser()

	c := &models.Category{
		Name:  "勉強",
		Color: "red",
	}

	user, _ := models.GetUser(2)

	// t := &models.Todos{}
	user.CreateTodo("プログラミング", c)

	// u, _ := models.GetUser(1)

	// fmt.Println(u)

	// u.DeleteUser()
}
