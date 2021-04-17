package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Compelo/compleo-api/types"
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

		//Eseguo la query
		queyRes, queyErr := db.Query("SELECT Descrizione FROM attivita")
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
	case "POST":
	case "PUT":
		//INSERISCI UN NUOVO LAVORO
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		var lavoroToAdd types.Activity

		//Popola l'oggetto
		lavoroToAdd.IDUtente, _ = strconv.ParseInt(r.FormValue("usrID"), 10, 64)

		lavoroToAdd.Testo = r.FormValue("testo")
		lavoroToAdd.Titolo = r.FormValue("titolo")
		lavoroToAdd.Tipo = r.FormValue("tipo")

		//TODO: CONTROLLA DATI

		//Esegui la query
		_, queyErr := db.Query("INSERT INTO `lavori`(`IDUtente`, `Tipo`, `Titolo`, `Testo`) VALUES ('" + fmt.Sprint(lavoroToAdd.IDUtente) + "','" + lavoroToAdd.Tipo + "','" + lavoroToAdd.Titolo + "','" + lavoroToAdd.Testo + "')")
		if queyErr != nil {
			fmt.Println(queyErr)
			w.Write([]byte(`{"message": "error"}`))
			fmt.Println("")
			return
		}
	case "DELETE":
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`NOT SUPPORTED`))
	}

	db.Close()

}
