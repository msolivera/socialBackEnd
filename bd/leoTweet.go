package bd

import (
	"context"
	"log"
	"time"

	"github.com/msolivera/socialTwittorBackEnd/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Funcion para leer los tweet de un perfil
//vamos a aprender a paginar en mongo
//devuelvo un slice con todos los tweet que voy a leer asi no tengo que
//consultar tantas veces cmo tweet hayan del usuario
func LeoTweet(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {

	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	//aca digo que y a trabajar sobre la tabla usuarios de la base de datos twittor
	coleccion := db.Collection("tweet")

	var resultados []*models.DevuelvoTweets

	//traigo los tweet del usuario con el id que traigo por parametro
	condicion := bson.M{

		"usuarioID": ID,
	}

	//opciones para filtrar y darle un comportamiento a mi consulta de bd
	opciones := options.Find()
	//creo objeto opciones y todo lo que setee ahora van a ser propiedades que intervienen
	//en el find
	//trabajaremos con el modo options en find
	opciones.SetLimit(20)                               //paginacion
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}}) //ordenamiento descendente
	opciones.SetSkip((pagina - 1) * 20)                 //esto es para que funcione la paginacion y mostrar de 20 en 20
	//ej> en la pagina 2 tengo que saltear los primeros 20 y en la pag 3 los primeros 40

	//creo puntero donde grabo resultados y puedo recorrer de a uno
	//para poder armar el resultado que mando a router
	cursor, err := coleccion.Find(contexto, condicion, opciones)

	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	for cursor.Next(context.TODO()) {
		var registro models.DevuelvoTweets
		//para cada elemento del cursos recorro y asigno a resultado
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}
	return resultados, true

}
