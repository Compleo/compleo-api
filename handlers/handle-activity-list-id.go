package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Compelo/compleo-api/types"
	_ "github.com/go-sql-driver/mysql"
)

func ActivityListPERIDHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a /activity/lid handle from ", r.RemoteAddr)
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
			fmt.Println("Errore")
			return
		}
		id := keys[0]

		w.WriteHeader(http.StatusOK)

		//Eseguo la query
		queyRes, queyErr := db.Query("SELECT * FROM `lavori` WHERE `IDUtente`='" + id + "'")
		if queyErr != nil {
			fmt.Println(queyErr)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		//Creo l'array da ritornare
		var array []types.Activity
		for queyRes.Next() {
			var g types.Activity
			queyRes.Scan(&g.ID, &g.IDUtente, &g.Tipo, &g.Titolo, &g.Testo, &g.UnitaMisura, &g.Prezzo, &g.Disponibilita)

			array = append(array, g)
		}

		//Create JSON
		j, _ := json.Marshal(array)
		w.Write([]byte(j))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`NOT SUPPORTED`))
	}

	db.Close()
}
