package types

type Activity struct {
	IDUtente    int
	IDAttivita  int
	Descrizione string
	Citta       City
	DataInizio  string
	DataFine    string
}
