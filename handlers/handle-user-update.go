package handlers

/*
   ***************************************
           Compleo Source Code
   ***************************************
   Programmer: Leonardo Baldazzi   (git -> @squirlyfoxy)
                                   (instagram -> @leonardobaldazzi_)

   Il seguente codice contiene gli handlers per la root /user/update

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

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a /user/update handle from ", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")

	db, sqlError := sql.Open("mysql", sqlVal)
	if sqlError != nil {
		panic(sqlError.Error())
	}

	switch r.Method {
	case "PUT":
		//AGGIORNA L'UTENTE
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		var toUpdateUser types.User

		//Popola l'oggetto
		toUpdateUser.ID, err = strconv.Atoi(r.FormValue("id"))
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(`{"message": "error"}`))
			return
		}

		toUpdateUser.Password = r.FormValue("password")
		toUpdateUser.Telefono = r.FormValue("telefono")
		toUpdateUser.EMail = r.FormValue("email")
		toUpdateUser.PartitaIVA = r.FormValue("piva")
		toUpdateUser.Bio = r.FormValue("bio")

		toUpdateUser.Sesso = r.FormValue("sesso")
		toUpdateUser.DataNascita = r.FormValue("dataNascita")

		//Esegui la query
		_, queyErr := db.Query("UPDATE `utente` SET `Telefono`='" + toUpdateUser.Telefono + "',`Bio`='" + toUpdateUser.Bio + "',`Email`='" + toUpdateUser.EMail + "',`IVA`='" + toUpdateUser.PartitaIVA + "',`Password`='" + toUpdateUser.Password + "', `Sesso`='" + toUpdateUser.Sesso + "', `DataNascita`='" + toUpdateUser.DataNascita + "' WHERE `ID` = '" + strconv.Itoa(toUpdateUser.ID) + "'")
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
