package main

import (
	"fmt"
	"timecontroller/app/controllers"
	"timecontroller/app/models"
)

func main() {
	controllers.StartMainServer()
	fmt.Println(models.Db)

}
