package main

import (
	"crud-app/client"
	"crud-app/database"
	"crud-app/use_cases"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the HTTP server on port 8090")
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
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/create", use_cases.CreatePerson).Methods("POST")
	router.HandleFunc("/get/{id}", use_cases.GetPersonByID).Methods("GET")
	router.HandleFunc("/update/{id}", use_cases.UpdatePersonByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", use_cases.DeletePersonByID).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8090", router))
}
