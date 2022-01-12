package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DevuelvoTweets struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuarioID string             `bson:"usuarioID" json:"usuarioID,omitempty"`
	Mensaje   string             `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha     time.Time          `bson:"fecha" json:"fecha,omitempty"`
}
