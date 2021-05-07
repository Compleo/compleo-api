package types

type Prenotazione struct {
	ID            int    `json:"id"`
	IDLavoro      int    `json:"idLavoro"`
	IDRichiedente int    `json:"idRichiedente"`
	Stato         string `json:"stato"`
	Scelta        string `json:"scelta"`
}
