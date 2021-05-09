package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Compelo/compleo-api/types"
	_ "github.com/go-sql-driver/mysql"
)

func ChatListAllRichiedenteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a /chat/get/richiedente handle from ", r.RemoteAddr)
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
		queyRes, queyErr := db.Query("SELECT * FROM `chat` WHERE `IDUtenteRichiedente`='" + id + "'")
		if queyErr != nil {
			fmt.Println(queyErr)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		var toReturn []types.Chat
		for queyRes.Next() {
			var g types.Chat
			scanErr := queyRes.Scan(&g.ID, &g.IDUtenteRichiedente, &g.IDUtenteDestinatario)
			if scanErr != nil {
				fmt.Println(scanErr)
				w.Write([]byte(`{"message": "error"}`))
				return
			}

			toReturn = append(toReturn, g)
		}

		//Create JSON
		j, _ := json.Marshal(toReturn)
		w.Write([]byte(j))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`NOT SUPPORTED`))
	}

	db.Close()
}
