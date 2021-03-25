package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//API functions:
//		work on 5051 port

var sqlVal = "root:@tcp(127.0.0.1:3306)/compleo"

type User struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Cognome   string `json:"cognome"`
	CF        string `json:"cf"`
	Indirizzo string `json:"indirizzo"`
	IDCitta   string `json:"idCitta"`
	Telefono  string `json:"telefono"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type POSTGotUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type City struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Provincia string `json:"provincia"`
}

type Province struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

func main() {
	fmt.Println(" -> Starting compleo-api version 0.1 server")
	fmt.Println(" -> Developed by Leonardo Baldazzi; Copyright (c) 2021 Compleo")
	fmt.Println(" -> Starting server on port 5051 and listening... ")

	http.HandleFunc("/", rootHanle)
	http.HandleFunc("/activity", activityHanle)
	http.HandleFunc("/user", userHandler)

	err := http.ListenAndServe(":5051", nil)
	checkError(err)
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func activityHanle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a /activity handle from ", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")

	db, sqlError := sql.Open("mysql", sqlVal)
	if sqlError != nil {
		panic(sqlError.Error())
	}

	switch r.Method {
	case "GET":
		//Controllo le variabili che mi sono state passate
		keys, err := r.URL.Query()["lst"] //Key lst sta per 'lista', se settata vuol dire che non devo listare
		if !err || len(keys[0]) < 1 {
			//Devo listare le attività

			//Eseguo la query
			queyRes, queyErr := db.Query("SELECT Descrizione FROM Attivita")
			if queyErr != nil {
				fmt.Println(queyErr)
				w.Write([]byte(`{"message": "error"}`))
				return
			}

			//Create array
			var array []string
			for queyRes.Next() {
				var g string
				queyRes.Scan(&g)

				array = append(array, g)
			}

			//Create JSON
			j, _ := json.Marshal(array)
			w.Write([]byte(j))

			return
		}

		lst := keys[0]

		if lst == "1" {
			w.Write([]byte(`TEST`)) //Farò qualcos'altro
		}

	case "POST":
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`NOT SUPPORTED`))
	}

}

func rootHanle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Someone has tried to acces to the root api, aborted")

	w.Write([]byte(`PERMISSION DENIED`))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a /user handle from ", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")

	db, sqlError := sql.Open("mysql", sqlVal)
	if sqlError != nil {
		panic(sqlError.Error())
	}

	switch r.Method {
	case "GET":
		//Posso passare per get solo l'username dell'utente, avrò in output: Nome, Cognome, Indirizzo, Telefono, Citta
		keys, err := r.URL.Query()["sr"]
		if !err || len(keys[0]) < 1 {
			w.Write([]byte(`{"message": "error"}`))
			return
		}
		usrName := keys[0]
		w.WriteHeader(http.StatusOK)

		//Execute query
		queyRes, queyErr := db.Query("SELECT Nome, Cognome, Indirizzo, IDCitta, Telefono FROM utente WHERE Account='" + usrName + "'")
		if queyErr != nil {
			fmt.Println(queyErr)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		//Create the object
		var userToRet User
		for queyRes.Next() {
			scanErr := queyRes.Scan(&userToRet.Nome, &userToRet.Cognome, &userToRet.Indirizzo, &userToRet.IDCitta, &userToRet.Telefono)
			if scanErr != nil {
				fmt.Println(scanErr)
				w.Write([]byte(`{"message": "error"}`))
				return
			}
		}

		userToRet.Username = usrName

		if userToRet.Nome != "" {
			//Create json
			j, jsonErr := json.Marshal(userToRet)
			if jsonErr != nil {
				fmt.Println(jsonErr)
				w.Write([]byte(`{"message": "error"}`))
				return
			}

			w.Write([]byte(j))
		} else {
			w.Write([]byte(`{"message": "error"}`))
		}
	case "POST":
		//Passo l'username e la password per avere tutte le informazioni di cui dispongo nel database
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		var pgu POSTGotUser
		json.Unmarshal([]byte(reqBody), &pgu)

		username := pgu.Username
		password := pgu.Password

		w.WriteHeader(http.StatusCreated)

		//Execute quey
		queyRes, queyErr := db.Query("SELECT * FROM utente WHERE Account='" + username + "' AND Password='" + password + "'")
		if queyErr != nil {
			fmt.Println(queyErr)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		//Create the object
		var userToRet User
		for queyRes.Next() {
			scanErr := queyRes.Scan(&userToRet.ID, &userToRet.Nome, &userToRet.Cognome, &userToRet.CF, &userToRet.Indirizzo, &userToRet.IDCitta, &userToRet.Telefono, &userToRet.Username, &userToRet.Password)
			if scanErr != nil {
				fmt.Println(scanErr)
				w.Write([]byte(`{"message": "error"}`))
				return
			}
		}

		if userToRet.Nome != "" {
			//Create json
			j, jsonErr := json.Marshal(userToRet)
			if jsonErr != nil {
				fmt.Println(jsonErr)
				w.Write([]byte(`{"message": "error"}`))
				return
			}

			w.Write([]byte(j))
		} else {
			w.Write([]byte(`{"message": "userNotFound"}`))
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`NOT SUPPORTED`))
	}

	db.Close()
}
