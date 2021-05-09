package types

type Chat struct {
	ID                   int   `json:"id"`
	IDUtenteRichiedente  int64 `json:"idUtenteRichiedente"`
	IDUtenteDestinatario int64 `json:"idUtenteDestinatario"`
}
