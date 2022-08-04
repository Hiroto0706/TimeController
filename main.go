package main

import (
	"timecontroller/app/models"
)

func main() {
	// fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = " ナランチャ"
	// u.Email = "test@test.com"

	// fmt.Println(u)
	// u.CreateUser()

	// cate, _ := models.GetCategory(4)
	// fmt.Println(cate.Name, cate.Color)
	// cate.Name = "健康系"
	// cate.Color = "orange"
	// fmt.Println(cate.Name, cate.Color)
	// cate.UpdateCategory()

	// cate.DeleteCategory()

	// user, _ := models.GetUser(1)
	// fmt.Println(user)

	// t := &models.Todo{}
	// user.CreateTodo("英語の勉強", c)

	// fmt.Println(t)

	// user, _ := models.GetUser(1)

	// fmt.Println(user)

	todo, _ := models.GetTodo(4)

	todo.DeleteTodo()

}
