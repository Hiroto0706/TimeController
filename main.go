package main

import (
	"fmt"
	"timecontroller/app/controllers"
	"timecontroller/config"
)

func main() {
	controllers.StartMainServer()

	// fmt.Println(models.Db)

	fmt.Println(config.Config.Static)

}
