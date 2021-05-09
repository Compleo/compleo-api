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

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a /chat handle from ", r.RemoteAddr)
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
		queyRes, queyErr := db.Query("SELECT * FROM `chat` WHERE `ID`='" + id + "'")
		if queyErr != nil {
			fmt.Println(queyErr)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		var toReturn types.Chat
		for queyRes.Next() {
			scanErr := queyRes.Scan(&toReturn.ID, &toReturn.IDUtenteRichiedente, &toReturn.IDUtenteDestinatario)
			if scanErr != nil {
				fmt.Println(scanErr)
				w.Write([]byte(`{"message": "error"}`))
				return
			}
		}

		//Create JSON
		j, _ := json.Marshal(toReturn)
		w.Write([]byte(j))
	case "PUT":
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		var chatToAdd types.Chat

		//Popola l'oggetto
		chatToAdd.IDUtenteDestinatario, _ = strconv.ParseInt(r.FormValue("idUtenteDestinatario"), 10, 64)
		chatToAdd.IDUtenteRichiedente, _ = strconv.ParseInt(r.FormValue("idUtenteRichiedente"), 10, 64)

		//Esegui la query
		_, queyErr := db.Query("INSERT INTO `chat`(`IDUtenteRichiedente`, `IDUtenteDestinatario`) VALUES ('" + fmt.Sprint(chatToAdd.IDUtenteRichiedente) + "','" + fmt.Sprint(chatToAdd.IDUtenteDestinatario) + "')")
		if queyErr != nil {
			fmt.Println(queyErr)
			w.Write([]byte(`{"message": "error"}`))
			fmt.Println("")
			return
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`NOT SUPPORTED`))
	}

	db.Close()
}
