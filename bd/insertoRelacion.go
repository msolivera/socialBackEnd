package bd

import (
	"context"
	"time"

	"github.com/msolivera/socialTwittorBackEnd/models"
)

//creacion de una relacion entre dos usuarios
func InsertoRelacion(t models.Relacion) (bool, error) {
	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	//aca digo que y a trabajar sobre la tabla usuarios de la base de datos twittor
	coleccion := db.Collection("relacion")

	_, err := coleccion.InsertOne(contexto, t)
	if err != nil {
		return false, err
	}
	return true, nil

}
