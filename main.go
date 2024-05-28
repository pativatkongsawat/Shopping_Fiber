package main

import (
	"go_fiber/config"
	"go_fiber/database"
)

func main() {

	config.Init()
	database.ConnectMySQL()

}
