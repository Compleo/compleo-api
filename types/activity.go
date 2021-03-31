package types

type Activity struct {
	IDUtente    int    `json:"idUtente"`
	IDAttivita  int    `json:"idAttivita"`
	Descrizione string `json:"descrizione"`
	Citta       City   `json:"citta"`
	DataInizio  string `json:"dataInizio"`
	DataFine    string `json:"dataFine"`
}
