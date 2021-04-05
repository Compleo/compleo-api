package handlers

/*
   ***************************************
           Compleo Source Code
   ***************************************
   Programmer: Leonardo Baldazzi   (git -> @squirlyfoxy)
                                   (instagram -> @leonardobaldazzi_)

   Il seguente codice contiene gli handlers per la root /recensione

   THE FOLLOWING SOURCE CODE IS CLOSED SOURCE
*/

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Compelo/compleo-api/types"
	_ "github.com/go-sql-driver/mysql"
)

func RecensioneHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a /recensione handle from ", r.RemoteAddr)
	db, sqlError := sql.Open("mysql", sqlVal)
	if sqlError != nil {
		panic(sqlError.Error())
	}

	switch r.Method {
	case "GET":
		//recID = IDRECENSITO, SE SETTATO RESTITUISCI TUTTE LE RECENSIONI DEL RECENSITO
		//redID = IDRECENSORE, SE SETTATO RESTITUISCI TUTTE LE RECENSIONI DEL RECENSORE
		keysC, errC := r.URL.Query()["recID"]
		if !errC || len(keysC[0]) < 1 {
			//NON ESISTE UNA KEY DI NOME recID, CONTROLLO SE ESISTE redID
			keysD, errD := r.URL.Query()["redID"]
			if !errD || len(keysD[0]) < 1 {
				w.Write([]byte(`{"message": "error"}`))
				return
			}

			//ESISTE redID, TODO: LISTA

			return
		}

		//recID := keysC[0]
		//TODO: LISTA

	case "PUT":
		//INSERISCI UNA NUOVA RECENSIONE
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		var recensioneToAdd types.Recensione

		//Popola l'oggetto
		recensioneToAdd.IDRecensito, _ = strconv.ParseInt(r.FormValue("recID"), 10, 64)
		recensioneToAdd.IDRecensore, _ = strconv.ParseInt(r.FormValue("redID"), 10, 64)

		recensioneToAdd.Testo = r.FormValue("testo")
		recensioneToAdd.Titolo = r.FormValue("titolo")
		recensioneToAdd.Valore, _ = strconv.ParseFloat(r.FormValue("voto"), 64)

		//TODO: CONTROLLA DATI

		//Esegui la query
		_, queyErr := db.Query("INSERT INTO `recensioni`(`IDRecensito`, `IDRecensore`, `Valore`, `Titolo`, `Testo`) VALUES ('" + fmt.Sprint(recensioneToAdd.IDRecensito) + "','" + fmt.Sprint(recensioneToAdd.IDRecensore) + "','" + fmt.Sprint(recensioneToAdd.Valore) + "','" + recensioneToAdd.Titolo + "','" + recensioneToAdd.Testo + "')")
		if queyErr != nil {
			fmt.Println(queyErr)
			w.Write([]byte(`{"message": "error"}`))
			fmt.Println("")
			return
		}
	case "DELETE":
		//ELIMINA UNA RECENSIONE (SAPENDO L'ID)
		keys, err := r.URL.Query()["id"]
		if !err || len(keys[0]) < 1 {
			w.Write([]byte(`{"message": "error"}`))
			return
		}
		recID := keys[0]
		w.WriteHeader(http.StatusOK)

		//Execute query
		_, queyErr := db.Query("DELETE FROM `recensioni` WHERE `ID`='" + recID + "'")
		if queyErr != nil {
			fmt.Println(queyErr)
			w.Write([]byte(`{"message": "error"}`))
			return
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`NOT SUPPORTED`))
	}

	db.Close()
}
