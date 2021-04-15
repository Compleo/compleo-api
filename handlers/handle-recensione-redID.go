package handlers

/*
   ***************************************
           Compleo Source Code
   ***************************************
   Programmer: Leonardo Baldazzi   (git -> @squirlyfoxy)
                                   (instagram -> @leonardobaldazzi_)

   Il seguente codice contiene gli handlers per la root /recensione/red

   THE FOLLOWING SOURCE CODE IS CLOSED SOURCE
*/

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Compelo/compleo-api/types"
	_ "github.com/go-sql-driver/mysql"
)

//ID RECENSORE

func RecensioneREDHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a /recensione/red handle from ", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")

	db, sqlError := sql.Open("mysql", sqlVal)
	if sqlError != nil {
		panic(sqlError.Error())
	}

	switch r.Method {
	case "GET":
		keys, err := r.URL.Query()["red"]
		if !err || len(keys[0]) < 1 {
			w.Write([]byte(`{"message": "error"}`))
			return
		}
		uID := keys[0]
		w.WriteHeader(http.StatusOK)

		//Eseguo la query
		queyRes, queyErr := db.Query("SELECT * FROM `recensioni` WHERE `IDRecensore`='" + uID + "'")
		if queyErr != nil {
			fmt.Println(queyErr)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		if queyErr != nil {
			fmt.Println(queyErr)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		var array []types.Recensione
		for queyRes.Next() {
			var g types.Recensione
			queyRes.Scan(&g.ID, &g.IDRecensito, &g.IDRecensore, &g.Valore, &g.Titolo, &g.Testo)

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
