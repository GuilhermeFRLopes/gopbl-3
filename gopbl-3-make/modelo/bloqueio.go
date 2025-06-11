package modelo

import "time"

type BloqueioPosto struct {
	ID           string    `bson:"_id"`
	PostoID      string    `bson:"postoId"`
	ClienteID    string    `bson:"clienteId"`
	DataBloqueio time.Time `bson:"dataBloqueio"`
	Expiracao    time.Time `bson:"expiracao"`
}