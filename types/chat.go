package types

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	TYPE_TEXT  = 0
	TYPE_IMAGE = 1
)

type Session struct {
	ID           int       `json:"id"`
	IdSender     int       `json:"idSender"`
	IdLavoratore int       `json:"idLavoratore"`
	IdLavoro     int       `json:"idLavoro"`
	CreatedTime  time.Time `json:"tempoCreazione"`
	Chat         []Chat    `json:"chat"`
}

type Chat struct {
	ID          int       `json:"id"`
	ContentType int       `json:"contentType"`
	Content     string    `json:"content"`
	CreatedTime time.Time `json:"tempoCreazione"`
}

func NewSession(idLavoratore int, idSender int, idLavoro int) Session {
	var s Session
	s.CreatedTime = time.Now()

	s.IdSender = idSender
	s.IdLavoratore = idLavoratore
	s.IdLavoro = idLavoro

	return s
}

func (s Session) NewChat(t int, content string) {
	var c Chat
	c.ContentType = t
	c.Content = content

	c.CreatedTime = time.Now()

	s.Chat = append(s.Chat, c)
}

func (s Session) SerializeSession() string {
	j, jsonErr := json.Marshal(s)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	return string(j)
}
