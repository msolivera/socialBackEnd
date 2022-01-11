package bd

import (
	"context"
	"time"

	"github.com/msolivera/socialTwittorBackEnd/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoTweet(t models.GraboTweet) (string, bool, error) {

	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	//aca digo que y a trabajar sobre la tabla usuarios de la base de datos twittor
	coleccion := db.Collection("tweet")

	//creamos el registro que vamos a insertar
	registro := bson.M{
		"usuarioID": t.UsuarioID,
		"mensaje":   t.Mensaje,
		"fecha":     t.Fecha,
	}
	//insert propiamente dicho
	result, err := coleccion.InsertOne(contexto, registro)
	if err != nil {
		return "", false, err
	}
	//extrae la clave del campo insertado (id)
	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.Hex(), true, nil

}
