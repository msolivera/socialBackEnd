package bd

import (
	"context"
	"time"

	"github.com/msolivera/socialTwittorBackEnd/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//funcion paramodificar el perfil del usuario
func ModificoRegistro(usu models.Usuario, ID string) (bool, error) {

	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	//aca digo que y a trabajar sobre la tabla usuarios de la base de datos twittor
	coleccion := db.Collection("usuarios")

	//make permite generar slices o mapas
	registro := make(map[string]interface{})
	if len(usu.Nombre) > 0 {
		registro["nombre"] = usu.Nombre
	}
	if len(usu.Apellidos) > 0 {
		registro["apellidos"] = usu.Apellidos
	}
	if len(usu.Avatar) > 0 {
		registro["avatar"] = usu.Avatar
	}
	if len(usu.Banner) > 0 {
		registro["banner"] = usu.Banner
	}
	if len(usu.Biografia) > 0 {
		registro["biografia"] = usu.Biografia
	}
	if len(usu.Ubicacion) > 0 {
		registro["ubicacion"] = usu.Ubicacion
	}
	if len(usu.SitioWeb) > 0 {
		registro["sitioWeb"] = usu.SitioWeb
	}
	registro["fechaNacimiento"] = usu.FechaNacimiento

	actualizacion := bson.M{

		"$set": registro,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)

	//creo el filtro para modificar la info del usuario con mismo id

	filtro := bson.M{
		"_id": bson.M{"$eq": objID},
	}
	//instruccion para saber si hubo o no un error
	_, err := coleccion.UpdateOne(contexto, filtro, actualizacion)

	if err != nil {
		return false, err
	}
	return true, nil
}
