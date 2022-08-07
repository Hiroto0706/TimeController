package main

import "timecontroller/app/controllers"

func main() {
	controllers.StartMainServer()

	// user, _ := models.GetUser(1)
	// fmt.Println(user)

	// if user.Password == models.Encrypt("password") {
	// 	fmt.Println(user.Password)
	// 	fmt.Println(models.Encrypt("password"))
	// }
}
