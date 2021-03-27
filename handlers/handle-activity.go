package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func ActivityHanle(w http.ResponseWriter, r *http.Request) {
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

	db.Close()

}
