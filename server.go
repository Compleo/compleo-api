package main

import (
	"fmt"
	"net/http"

	handlers "github.com/Compelo/compleo-api/handlers"

	_ "github.com/go-sql-driver/mysql"
)

//API functions:
//		work on port 5051

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println(" -> Starting compleo-api server version 0.4.7")
	fmt.Println(" -> Developed by Compleo' Developers; Copyright (c) 2021 Compleo")
	fmt.Println(" -> Starting server on port 5051 and listening... ")

	initHandlers()

	err := http.ListenAndServe(":5051", nil)
	checkError(err)
}

func initHandlers() {
	go http.HandleFunc("/", rootHanle)
	go http.HandleFunc("/activity", handlers.ActivityHanle)
	go http.HandleFunc("/recensione", handlers.RecensioneHandler)
	go http.HandleFunc("/user", handlers.UserHandler)
	go http.HandleFunc("/user/update", handlers.UpdateUserHandler)
}

func rootHanle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Someone has tried to acces to the root api, aborted")

	w.Write([]byte(`PERMISSION DENIED`))
}
