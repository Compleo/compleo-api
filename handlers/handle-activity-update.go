package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Compelo/compleo-api/types"
	_ "github.com/go-sql-driver/mysql"
)

func ActivityUpdateHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a /activity/update handle from ", r.RemoteAddr)
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

		var toUpdateActivity types.Activity

		//Popola l'oggetto
		toUpdateActivity.ID, err = strconv.ParseInt(r.FormValue("id"), 10, 64)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		toUpdateActivity.Testo = r.FormValue("testo")
		toUpdateActivity.Tipo = r.FormValue("tipo")
		toUpdateActivity.Titolo = r.FormValue("titolo")

		_, queyErr := db.Query("UPDATE `lavori` SET `Tipo`='" + toUpdateActivity.Tipo + "',`Titolo`='" + toUpdateActivity.Titolo + "',`Testo`='" + toUpdateActivity.Testo + "' WHERE `ID` = '" + strconv.Itoa(int(toUpdateActivity.ID)) + "'")
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
