package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
		//Posso passare per get solo l'username dell'utente, avrò in output: Nome, Cognome, Indirizzo, Telefono, Citta
		keys, err := r.URL.Query()["sr"]
		if !err || len(keys[0]) < 1 {
			w.Write([]byte(`{"message": "error"}`))
			return
		}
		usrName := keys[0]
		w.WriteHeader(http.StatusOK)

		//Execute query
		queyRes, queyErr := db.Query("SELECT Nome, Cognome, Indirizzo, Citta, Provincia, Regione, Telefono FROM utente WHERE Account='" + usrName + "'")
		if queyErr != nil {
			fmt.Println(queyErr)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		//Create the object
		var userToRet types.User
		for queyRes.Next() {
			scanErr := queyRes.Scan(&userToRet.Nome, &userToRet.Cognome, &userToRet.Indirizzo, &userToRet.Citta, &userToRet.Regione, &userToRet.Provincia, &userToRet.Telefono)
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
		for queyRes.Next() {
			scanErr := queyRes.Scan(&userToRet.ID, &userToRet.Nome, &userToRet.Cognome, &userToRet.CF, &userToRet.Indirizzo, &userToRet.Citta, &userToRet.Regione, &userToRet.Provincia, &userToRet.Telefono, &userToRet.Username, &userToRet.Password)
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
		toRegisterUser.Nome = r.FormValue("nome")
		toRegisterUser.Cognome = r.FormValue("cognome")
		toRegisterUser.CF = r.FormValue("cf")

		//Provenienza utente
		toRegisterUser.Indirizzo = r.FormValue("indirizzo")
		toRegisterUser.Citta = r.FormValue("citta")
		toRegisterUser.Provincia = r.FormValue("provincia")
		toRegisterUser.Regione = r.FormValue("regione")

		toRegisterUser.Telefono = r.FormValue("telefono")
		toRegisterUser.Username = r.FormValue("username")
		toRegisterUser.Password = r.FormValue("password")

		//Esegui la query
		_, queyErr := db.Query("INSERT INTO `utente`(`Nome`, `Cognome`, `CF`, `Indirizzo`, `Citta`, `Regione`, `Provincia`, `Telefono`, `Username`, `Password`) VALUES ('" + toRegisterUser.Nome + "', '" + toRegisterUser.Cognome + "', '" + toRegisterUser.CF + "', '" + toRegisterUser.Indirizzo + "', '" + toRegisterUser.Citta + "', '" + toRegisterUser.Regione + "', '" + toRegisterUser.Provincia + "', '" + toRegisterUser.Telefono + "', '" + toRegisterUser.Username + "', '" + toRegisterUser.Password + "')")
		if queyErr != nil {
			fmt.Println(queyErr)
			w.Write([]byte(`{"message": "error"}`))
			fmt.Println("INSERT INTO `utente`(`Nome`, `Cognome`, `CF`, `Indirizzo`, `Citta`, `Regione`, `Provincia`, `Telefono`, `Username`, `Password`) VALUES ('" + toRegisterUser.Nome + "', '" + toRegisterUser.Cognome + "', '" + toRegisterUser.CF + "', '" + toRegisterUser.Indirizzo + "', '" + toRegisterUser.Citta + "', '" + toRegisterUser.Regione + "', '" + toRegisterUser.Provincia + "', '" + toRegisterUser.Telefono + "', '" + toRegisterUser.Username + "', '" + toRegisterUser.Password + "')")
			return
		}

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`NOT SUPPORTED`))
	}

	db.Close()
}