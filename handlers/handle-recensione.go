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
	"fmt"
	"net/http"
)

func RecensioneHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a /recensione handle from ", r.RemoteAddr)
}
