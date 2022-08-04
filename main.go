package main

import (
	"fmt"
	"log"
	"timecontroller/config"
)

func main() {
	fmt.Println(config.Config.SQLDriver, config.Config.Port, config.Config.DbName, config.Config.LogFile)

	log.Println("this is a test")
}
