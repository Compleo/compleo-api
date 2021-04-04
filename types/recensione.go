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
	IDRecensito int     `json:"idRecensito"`
	IDRecensore int     `json:"idRecensore"`
	Valore      float32 `json:"valore"`
	Testo       string  `json:"testo"`
}
