package main

import (
	"fmt"
	"golang-pagination/src/config"
	"golang-pagination/src/models"
	"golang-pagination/src/routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	configDB, err := gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status: ", err)
	}

	defer configDB.Close()
	configDB.AutoMigrate(&models.User{})
	r := routes.SetupRouter()

	r.Run()
}
