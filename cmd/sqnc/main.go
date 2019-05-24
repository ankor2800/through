package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/doug-martin/goqu.v5"
	"log"
	"net/http"
	"through/controllers"
	"through/env"
	"through/store/mysql"
)

func main() {

	db, err := makeMysql()

	if err != nil {
		fmt.Printf("Failed to mysql connection: %v\n", err)
	}

	orders := mysql.NewOrderRepo(db)
	router := mux.NewRouter()
	controller := controllers.Controller{
		Order: orders,
	}
	router.HandleFunc("/add/through", controller.AddSequence()).Methods("POST")
	router.HandleFunc("/add/auto", controller.AddAutoincrement()).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func makeMysql() (*goqu.Database, error) {
	return mysql.NewConnection(&mysql.ConnectionOptions{
		Host:  env.Config.DBHost,
		Port:  env.Config.DBPort,
		Name:  env.Config.DBDatabase,
		User:  env.Config.DBUser,
		Pass:  env.Config.DBPassword,
		Debug: env.Config.DBDebug,
	})
}
