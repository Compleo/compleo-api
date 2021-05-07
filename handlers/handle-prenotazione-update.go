package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Compelo/compleo-api/types"
	_ "github.com/go-sql-driver/mysql"
)

func PrenotazioneUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a /prenotazione handle from ", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")

	db, sqlError := sql.Open("mysql", sqlVal)
	if sqlError != nil {
		panic(sqlError.Error())
	}

	switch r.Method {
	case "PUT":
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		var toUpdatePrenotazione types.Prenotazione

		//Popola l'oggetto
		toUpdatePrenotazione.ID, err = strconv.Atoi(r.FormValue("id"))
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		toUpdatePrenotazione.Stato = r.FormValue("stato")

		//Esegui la query
		_, queyErr := db.Query("UPDATE `prenotazioni` SET `Scelta`='" + toUpdatePrenotazione.Scelta + "' WHERE `ID`='" + fmt.Sprint(toUpdatePrenotazione.ID) + "'")
		if queyErr != nil {
			fmt.Println(err)
			w.Write([]byte(`{"message": "error"}`))
			return
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`NOT SUPPORTED`))
	}

	db.Close()
}
