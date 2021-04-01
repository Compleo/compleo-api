package main

import (
	"fmt"
	"net/http"

	handlers "github.com/Compelo/compleo-api/handlers"

	_ "github.com/go-sql-driver/mysql"
)

//API functions:
//		work on 5051 port

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println(" -> Starting compleo-api version 0.4 server")
	fmt.Println(" -> Developed by Compleo' Developers; Copyright (c) 2021 Compleo")
	fmt.Println(" -> Starting server on port 5051 and listening... ")

	initHandlers()

	err := http.ListenAndServe(":5051", nil)
	checkError(err)
}

func initHandlers() {
	http.HandleFunc("/", rootHanle)
	http.HandleFunc("/activity", handlers.ActivityHanle)
	http.HandleFunc("/user", handlers.UserHandler)
}

func rootHanle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Someone has tried to acces to the root api, aborted")

	w.Write([]byte(`PERMISSION DENIED`))
}
