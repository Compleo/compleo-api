package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Compelo/compleo-api/types"
	_ "github.com/go-sql-driver/mysql"
)

func GetByIDHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a /user/getByID handle from ", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")

	db, sqlError := sql.Open("mysql", sqlVal)
	if sqlError != nil {
		panic(sqlError.Error())
	}

	switch r.Method {
	case "GET":
		//Posso passare per get solo l'username dell'utente, avrò in output: Nome, Cognome, Indirizzo, Citta, Provincia, Regione, Telefono
		keys, err := r.URL.Query()["id"]
		if !err || len(keys[0]) < 1 {
			w.Write([]byte(`{"message": "error"}`))
			return
		}
		usrID := keys[0]
		w.WriteHeader(http.StatusOK)

		//Execute query
		queyRes, queyErr := db.Query("SELECT ID, Nome, Cognome, Indirizzo, Citta, Provincia, Regione, Telefono, Bio, Email, Livello FROM utente WHERE ID='" + usrID + "'")
		if queyErr != nil {
			fmt.Println(queyErr)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		//Create the object
		var userToRet types.User
		var cittaToRet types.City
		userToRet.Citta = cittaToRet
		for queyRes.Next() {
			scanErr := queyRes.Scan(&userToRet.ID, &userToRet.Nome, &userToRet.Cognome, &userToRet.Indirizzo, &userToRet.Citta.Nome, &userToRet.Citta.Provincia, &userToRet.Citta.Regione, &userToRet.Telefono, &userToRet.Bio, &userToRet.EMail, &userToRet.Livello)

			userToRet.Username = strings.ToLower(userToRet.Nome) + "." + strings.ToLower(userToRet.Cognome)

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
			w.Write([]byte(`{"message": "error"}`))
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`NOT SUPPORTED`))
	}

	db.Close()
}
