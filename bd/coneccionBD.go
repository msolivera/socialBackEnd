package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//archivo para generar la conexion a la BD
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://molivera:52304220@twittor.nyi9h.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

//Funcion que genera y verifica la conexion
func ConectarBD() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		//si hubo error en conexion
		log.Fatal(err.Error())
		return client
	}
	//ping para saber si la BD esta arriba
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		//hubo error en la BD
		log.Fatal(err.Error())
		return client

	}
	log.Println("Conexion exitosa con la BD")
	//conexion valida
	return client
}

func ChequeoConexion() int {

	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
