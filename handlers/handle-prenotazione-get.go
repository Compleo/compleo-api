package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Compelo/compleo-api/types"
	_ "github.com/go-sql-driver/mysql"
)

func PrenotazioneGetDaIDHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a /prenotazione/get handle from ", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")

	db, sqlError := sql.Open("mysql", sqlVal)
	if sqlError != nil {
		panic(sqlError.Error())
	}

	switch r.Method {
	case "GET":
		keys, err := r.URL.Query()["id"]
		if !err || len(keys[0]) < 1 {
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		id := keys[0]
		w.WriteHeader(http.StatusOK)

		//Execute query
		queyRes, queyErr := db.Query("SELECT * FROM prenotazioni WHERE ID='" + id + "'")
		if queyErr != nil {
			fmt.Println(queyErr)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		var toReturnPrenotazione types.Prenotazione

		for queyRes.Next() {
			scanErr := queyRes.Scan(&toReturnPrenotazione.ID, &toReturnPrenotazione.IDLavoro, &toReturnPrenotazione.IDRichiedente, &toReturnPrenotazione.Stato, &toReturnPrenotazione.Scelta)
			if scanErr != nil {
				fmt.Println(scanErr)
				w.Write([]byte(`{"message": "error"}`))
				return
			}
		}

		//Create json
		j, jsonErr := json.Marshal(toReturnPrenotazione)
		if jsonErr != nil {
			fmt.Println(jsonErr)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		w.Write([]byte(j))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`NOT SUPPORTED`))
	}

	db.Close()
}
