package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/msolivera/socialTwittorBackEnd/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ConsultoRelacion(t models.Relacion) (bool, error) {

	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	//aca digo que y a trabajar sobre la tabla usuarios de la base de datos twittor
	coleccion := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion
	fmt.Println(resultado)
	err := coleccion.FindOne(contexto, condicion).Decode(&resultado)

	if err != nil {
		return false, err
	}
	return true, nil

}
