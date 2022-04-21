package main

import (
	"crud-app/client"
	"crud-app/database"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "root",
			DB:         "go",
		}

	connectionString := database.GetConnectionString(config)
	err := client.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
}
