package bd

import (
	"context"
	"time"

	"github.com/msolivera/socialTwittorBackEnd/models"
	"go.mongodb.org/mongo-driver/bson"
)

//recibe un mail por parametro y comprueba si esta en la BD
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	//aca digo que y a trabajar sobre la tabla usuarios de la base de datos twittor
	coleccion := db.Collection("usuarios")

	//bson.M trae un mapa tipo clave valor
	//en bson trabaja la bd de mongo
	condicion := bson.M{"email": email}

	var resultado models.Usuario

	//voy a la coleccion y busco la condicion en el contexto, y lo decodifico en un json
	//el json se lo asigno al resultado que es un usuario
	err := coleccion.FindOne(contexto, condicion).Decode(&resultado)

	//recibe el id y lo convierte en hexadecimal string
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID

}
