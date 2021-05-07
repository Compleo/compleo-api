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

func PrenotazioneHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a /prenotazione handle from ", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")

	db, sqlError := sql.Open("mysql", sqlVal)
	if sqlError != nil {
		panic(sqlError.Error())
	}

	switch r.Method {
	case "GET":
		//Prendi le prenotazioni di un utente
		keys, err := r.URL.Query()["idUtente"]
		if !err || len(keys[0]) < 1 {
			w.Write([]byte(`{"message": "error"}`))
			fmt.Println("Errore")
			return
		}
		id := keys[0]

		w.WriteHeader(http.StatusOK)

		//Eseguo la query
		queyRes, queyErr := db.Query("SELECT * FROM `prenotazioni` WHERE `IDRichiedente`='" + id + "'")
		if queyErr != nil {
			fmt.Println(queyErr)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		//Creo l'array da ritornare
		var array []types.Prenotazione
		for queyRes.Next() {
			var g types.Prenotazione
			queyRes.Scan(&g.ID, &g.IDLavoro, &g.IDRichiedente, &g.Stato, &g.Scelta)

			array = append(array, g)
		}

		//Create JSON
		j, _ := json.Marshal(array)
		w.Write([]byte(j))
	case "PUT":
		//Inserisci una nuova prenotazione
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		var toAddPrenotazione types.Prenotazione

		//Popola l'oggetto
		toAddPrenotazione.IDLavoro, _ = strconv.Atoi(r.FormValue("idLavoro"))
		toAddPrenotazione.IDRichiedente, _ = strconv.Atoi(r.FormValue("idRichiedente"))
		toAddPrenotazione.Scelta = r.FormValue("scelta")
		toAddPrenotazione.Stato = "Richiesto"

		//Esegui la query
		_, queyErr := db.Query("INSERT INTO `prenotazioni`(`IDLavoro`, `IDRichiedente`, `Stato`, `Scelta`) VALUES ('" + fmt.Sprint(toAddPrenotazione.IDLavoro) + "','" + fmt.Sprint(toAddPrenotazione.IDRichiedente) + "','" + toAddPrenotazione.Stato + "','" + toAddPrenotazione.Scelta + "')")
		if queyErr != nil {
			fmt.Println(queyErr)
			w.Write([]byte(`{"message": "error"}`))
			fmt.Println("")
			return
		}

	case "DELETE":
		//Elimina la prenotazione
		keys, err := r.URL.Query()["id"]
		if !err || len(keys[0]) < 1 {
			w.Write([]byte(`{"message": "error"}`))
			return
		}
		userID := keys[0]
		w.WriteHeader(http.StatusOK)

		//Execute query
		_, queyErr := db.Query("DELETE FROM `prenotazioni` WHERE `ID`='" + userID + "'")
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
