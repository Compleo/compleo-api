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
	fmt.Println(" -> Starting compleo-api server version 0.5.0")
	fmt.Println(" -> Developed by Compleo' Developers; Copyright (c) 2021 Compleo")
	fmt.Println(" -> Starting server on port 5051 and listening... ")

	initHandlers()

	err := http.ListenAndServe(":5051", nil)
	checkError(err)
}

func initHandlers() {
	//Root
	go http.HandleFunc("/", rootHanle)

	//Attività
	go http.HandleFunc("/activity", handlers.ActivityHanle)
	go http.HandleFunc("/activity/lid", handlers.ActivityListPERIDHandle)
	go http.HandleFunc("/activity/listqual", handlers.ActivityListPerQualificheHandle)
	go http.HandleFunc("/activity/listall", handlers.ActivityListAllHanle)
	go http.HandleFunc("/activity/getbyid", handlers.ActivityGetFromIDHanle)
	go http.HandleFunc("/activity/update", handlers.ActivityUpdateHandle)

	//Recensioni
	go http.HandleFunc("/recensione", handlers.RecensioneHandler)
	go http.HandleFunc("/recensione/rec", handlers.RecensioneRECHandler)
	go http.HandleFunc("/recensione/red", handlers.RecensioneREDHandler)
	go http.HandleFunc("/recensione/get", handlers.RecensioneGetIDHandler)

	//Prenotazioni
	go http.HandleFunc("/prenotazione", handlers.PrenotazioneHandler)
	go http.HandleFunc("/prenotazione/get", handlers.PrenotazioneGetDaIDHandler)
	go http.HandleFunc("/prenotazione/update", handlers.PrenotazioneUpdateHandler)
	go http.HandleFunc("/prenotazione/get/all", handlers.PrenotazioneGetAllFromIdLavoroHandler)

	//Chat
	go http.HandleFunc("/chat", handlers.ChatHandler)
	go http.HandleFunc("/chat/get/destinatario", handlers.ChatListAllDestinatarioHandler)
	go http.HandleFunc("/chat/get/richiedente", handlers.ChatListAllRichiedenteHandler)

	//Utenti
	go http.HandleFunc("/user", handlers.UserHandler)
	go http.HandleFunc("/user/update", handlers.UpdateUserHandler)
	go http.HandleFunc("/user/getByID", handlers.GetByIDHandler)
}

func rootHanle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Someone has tried to acces to the root api, aborted")

	w.Write([]byte(`PERMISSION DENIED`))
}
