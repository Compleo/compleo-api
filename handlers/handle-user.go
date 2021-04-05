package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	types "github.com/Compelo/compleo-api/types"
	_ "github.com/go-sql-driver/mysql"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a /user handle from ", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")

	db, sqlError := sql.Open("mysql", sqlVal)
	if sqlError != nil {
		panic(sqlError.Error())
	}

	switch r.Method {
	case "GET":
		//Posso passare per get solo l'username dell'utente, avrò in output: Nome, Cognome, Indirizzo, Citta, Provincia, Regione, Telefono
		keys, err := r.URL.Query()["sr"]
		if !err || len(keys[0]) < 1 {
			w.Write([]byte(`{"message": "error"}`))
			return
		}
		usrName := keys[0]
		w.WriteHeader(http.StatusOK)

		//Execute query
		queyRes, queyErr := db.Query("SELECT ID, Nome, Cognome, Indirizzo, Citta, Provincia, Regione, Telefono, Bio, Email, Livello FROM utente WHERE Username='" + usrName + "'")
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
			if scanErr != nil {
				fmt.Println(scanErr)
				w.Write([]byte(`{"message": "error"}`))
				return
			}
		}

		userToRet.Username = usrName

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
	case "POST":
		//Passo l'username e la password per avere tutte le informazioni di cui dispongo nel database
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		var pgu types.POSTGotUser
		json.Unmarshal([]byte(reqBody), &pgu)

		username := pgu.Username
		password := pgu.Password

		if pgu.CheckUser() {
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		w.WriteHeader(http.StatusCreated)

		//Execute quey
		queyRes, queyErr := db.Query("SELECT * FROM utente WHERE Username='" + username + "' AND Password='" + password + "'")
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
			scanErr := queyRes.Scan(&userToRet.ID, &userToRet.Nome, &userToRet.Cognome, &userToRet.CF, &userToRet.Indirizzo, &userToRet.Citta.Nome, &userToRet.Citta.Regione, &userToRet.Citta.Provincia, &userToRet.Telefono, &userToRet.Bio, &userToRet.EMail, &userToRet.Username, &userToRet.Livello, &userToRet.PartitaIVA, &userToRet.Password)
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
			w.Write([]byte(`{"message": "userNotFound"}`))
		}
	case "PUT":
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		var toRegisterUser types.User
		//Popola l'oggetto

		//Informazione utente
		toRegisterUser.Nome = r.FormValue("nome")
		toRegisterUser.Cognome = r.FormValue("cognome")
		toRegisterUser.Telefono = r.FormValue("telefono")
		toRegisterUser.CF = r.FormValue("cf")

		//Livello utente
		toRegisterUser.Livello = r.FormValue("lvl")

		//Controllo che sia un livello accettato
		if toRegisterUser.Livello != "Base" && toRegisterUser.Livello != "Completo" {
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		//Partita iva
		if iva := r.FormValue("piva"); iva != "" {
			toRegisterUser.PartitaIVA = iva
		}

		//Posizione geografica
		toRegisterUser.Indirizzo = r.FormValue("indirizzo")

		var citta types.City
		citta.Nome = r.FormValue("citta")
		citta.Regione = r.FormValue("regione")
		citta.Provincia = r.FormValue("provincia")

		toRegisterUser.Citta = citta

		//Utente
		toRegisterUser.EMail = r.FormValue("email")
		toRegisterUser.Password = r.FormValue("password")

		//Creo il nome utente partendo dal nome e cognome
		lowerNome := strings.ToLower(toRegisterUser.Nome)
		lowerCognome := strings.ToLower(toRegisterUser.Cognome)

		toRegisterUser.Username = lowerNome + "." + lowerCognome

		//TODO: CONTROLLA SE UN UTENTE CON QUESTO USERNAME E' GIA' REGISTRATO, SE SI MODIFICO L'USERNAME
		/*
			//Controlla che i dati siano corretti
			if !toRegisterUser.CheckUser() {
				w.Write([]byte(`{"message": "error"}`))
				fmt.Println("ERRORE")
				fmt.Println(toRegisterUser.Username)
				return
			}*/

		//Esegui la query
		_, queyErr := db.Query("INSERT INTO `utente`(`Nome`, `Cognome`, `CF`, `Indirizzo`, `Citta`, `Regione`, `Provincia`, `Telefono`, `Email`, `Username`, `Livello`, `IVA`, `Password`) VALUES ('" + toRegisterUser.Nome + "', '" + toRegisterUser.Cognome + "', '" + toRegisterUser.CF + "', '" + toRegisterUser.Indirizzo + "', '" + toRegisterUser.Citta.Nome + "', '" + toRegisterUser.Citta.Regione + "', '" + toRegisterUser.Citta.Provincia + "', '" + toRegisterUser.Telefono + "', '" + toRegisterUser.EMail + "', '" + toRegisterUser.Username + "', '" + toRegisterUser.Livello + "', '" + toRegisterUser.PartitaIVA + "', '" + toRegisterUser.Password + "')")
		if queyErr != nil {
			fmt.Println(queyErr)
			w.Write([]byte(`{"message": "error"}`))
			fmt.Println("INSERT INTO `utente`(`Nome`, `Cognome`, `CF`, `Indirizzo`, `Citta`, `Regione`, `Provincia`, `Telefono`, `Email`, `Username`, `Password`) VALUES ('" + toRegisterUser.Nome + "', '" + toRegisterUser.Cognome + "', '" + toRegisterUser.CF + "', '" + toRegisterUser.Indirizzo + "', '" + toRegisterUser.Citta.Nome + "', '" + toRegisterUser.Citta.Regione + "', '" + toRegisterUser.Citta.Provincia + "', '" + toRegisterUser.Telefono + "', '" + toRegisterUser.EMail + "', '" + toRegisterUser.Username + "', '" + toRegisterUser.Password + "')")
			return
		}
	case "DELETE":
		keys, err := r.URL.Query()["id"]
		if !err || len(keys[0]) < 1 {
			w.Write([]byte(`{"message": "error"}`))
			return
		}
		userID := keys[0]
		w.WriteHeader(http.StatusOK)

		//Execute query
		_, queyErr := db.Query("DELETE FROM `utente` WHERE `ID`='" + userID + "'")
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
