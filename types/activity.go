package types

type Activity struct {
	ID            int64   `json:"id"`
	IDUtente      int64   `json:"idUtente"`
	Tipo          string  `json:"tipo"`
	Titolo        string  `json:"titolo"`
	Testo         string  `json:"testo"`
	UnitaMisura   string  `json:"unitaMisura"`
	Prezzo        float64 `json:"prezzo"`
	Disponibilita string  `json:"disponibilita"`
}
