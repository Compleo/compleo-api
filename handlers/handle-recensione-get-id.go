package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Compelo/compleo-api/types"
)

func RecensioneGetIDHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a /recensione/get handle from ", r.RemoteAddr)
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
		queyRes, queyErr := db.Query("SELECT * FROM recensioni WHERE ID='" + id + "'")
		if queyErr != nil {
			fmt.Println(queyErr)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		var recensioneToRet types.Recensione
		for queyRes.Next() {
			scanErr := queyRes.Scan(&recensioneToRet.ID, &recensioneToRet.IDRecensito, &recensioneToRet.IDRecensore, &recensioneToRet.Valore, &recensioneToRet.Titolo, &recensioneToRet.Testo)
			if scanErr != nil {
				fmt.Println(scanErr)
				w.Write([]byte(`{"message": "error"}`))
				return
			}
		}

		//Create json
		j, jsonErr := json.Marshal(recensioneToRet)
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
