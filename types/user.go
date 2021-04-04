package types

type User struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Cognome   string `json:"cognome"`
	CF        string `json:"cf"`
	Indirizzo string `json:"indirizzo"`

	Citta City `json:"citta"`

	Telefono   string `json:"telefono"`
	Bio        string `json:"bio"`
	EMail      string `json:"email"`
	Username   string `json:"username"`
	Livello    string `json:"livello"`
	PartitaIVA string `json:"piva"`
	Password   string `json:"password"`
}

type POSTGotUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u POSTGotUser) CheckUser() bool {
	return u.Username == "" || u.Password == ""
}

func (u User) CheckUser() bool {
	return u.Nome == "" || u.Cognome == "" || u.CF == "" || u.Indirizzo == "" || u.Citta.Regione == "" || u.Citta.Nome == "" || u.Citta.Provincia == "" || u.Telefono == "" || u.EMail == "" || u.Username == "" || u.Password == ""
}
