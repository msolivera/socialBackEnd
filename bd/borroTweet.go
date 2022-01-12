package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BorroTweet(IDTweet string, IDUsuario string) error {
	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	//aca digo que y a trabajar sobre la tabla usuarios de la base de datos twittor
	coleccion := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(IDTweet)

	condicion := bson.M{
		"_id":       objID,
		"usuarioID": IDUsuario,
	}

	_, err := coleccion.DeleteOne(contexto, condicion)
	return err
}
