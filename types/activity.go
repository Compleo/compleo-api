package types

type Activity struct {
	ID       int    `json:"id"`
	IDUtente int    `json:"idUtente"`
	Tipo     string `json:"tipo"`
	Titolo   string `json:"titolo"`
	Testo    string `json:"testo"`
}
