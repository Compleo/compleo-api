package types

type User struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Cognome   string `json:"cognome"`
	CF        string `json:"cf"`
	Indirizzo string `json:"indirizzo"`
	Regione   string `json:"regione"`
	Citta     string `json:"citta"`
	Provincia string `json:"provincia"`
	Telefono  string `json:"telefono"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type POSTGotUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
