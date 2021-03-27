package types

type City struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Provincia string `json:"provincia"`
}

type Province struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}
