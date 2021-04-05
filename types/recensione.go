package types

/*
   ***************************************
           Compleo Source Code
   ***************************************
   Programmer: Leonardo Baldazzi   (git -> @squirlyfoxy)
                                   (instagram -> @leonardobaldazzi_)

   Il seguente codice rappresenta un oggetto di ripo "Recensione"

   THE FOLLOWING SOURCE CODE IS CLOSED SOURCE
*/

type Recensione struct {
	ID          int     `json:"id"`
	IDRecensito int64   `json:"idRecensito"`
	IDRecensore int64   `json:"idRecensore"`
	Valore      float64 `json:"valore"`
	Titolo      string  `json:"titolo"`
	Testo       string  `json:"testo"`
}
