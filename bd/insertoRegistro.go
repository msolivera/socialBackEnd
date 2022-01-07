package bd

import (
	"context"
	"time"

	"github.com/msolivera/socialTwittorBackEnd/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//aca inserto efectivamente los registros de usuario en la bd
func InsertoRegistro(usu models.Usuario) (string, bool, error) {

	//creacion de un contexto
	//asegurarme de que la bd no quede colgada ni naa
	//quiero un contexto especial
	//El context.backgroun es el que traigo de la bd, es como que le agrego un timeout de 15 segundos
	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//instruccion que seteo al comienzo pero se ejecuta al final de la funcion
	defer cancel()
	//cancel cancela el contexto.withtimeout, asi no ocupa espacio inecesario
	db := MongoCN.Database("twittor")
	//aca digo que y a trabajar sobre la tabla usuarios de la base de datos twittor
	coleccion := db.Collection("usuarios")

	//esto para encriptar la pass asi no se gruarda como texto claro
	usu.Password, _ = EncriptarPassword(usu.Password)

	result, err := coleccion.InsertOne(contexto, usu)
	if err != nil {
		return "", false, err
	}

	//esto no lo voy a usar pero es para saber como obtener el id del resultado
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
