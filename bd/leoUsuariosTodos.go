package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/msolivera/socialTwittorBackEnd/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*lee los usuarios registrados en el sistema*/
/* id = mi id
page = cantidad de registros a mostrar por pagina
search = criterio de busqueda
tipo = tipo de busqueda
devuelve un slice e usuarios y un bool */
func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {

	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	//aca digo que y a trabajar sobre la tabla usuarios de la base de datos twittor
	coleccion := db.Collection("usuarios")
	//slice donde voy a guardar los resultados
	var resultados []*models.Usuario

	findOption := options.Find()
	findOption.SetSkip((page - 1) * 20)
	findOption.SetLimit(20)

	condicion := bson.M{
		//busca sin importar mayusculas y minusculas
		"nombre": bson.M{"$regex": `(?i)` + search},
	}
	cursor, err := coleccion.Find(contexto, condicion, findOption)

	if err != nil {
		fmt.Println(err.Error())
		return resultados, false
	}

	var encontrado, incluir bool
	//empiezo a recorrer el slice de usuarios
	for cursor.Next(contexto) {
		//creo un usuario y le asigno el ususario del cursor
		var usu models.Usuario
		err := cursor.Decode(&usu)
		if err != nil {
			fmt.Println(err.Error())
			return resultados, false
		}
		var relacion models.Relacion
		//si no hay error asigno al usu creado las variables del cursor
		relacion.UsuarioID = ID
		relacion.UsuarioRelacionID = usu.ID.Hex()
		//seteo la variable en false
		incluir = false

		//consulto si ambos id tienen relacion
		encontrado, err = ConsultoRelacion(relacion)
		//si el tipo es nuevo (todos los usuarios) y no se relaciona conmigo
		if tipo == "new" && encontrado == false {
			incluir = true
			fmt.Println("new")
		}
		//si el tipo es los que yo sigo y ya tengo relacion
		if tipo == "follow" && encontrado == true {
			incluir = true
			fmt.Println("follow")
		}
		//si yo me sigo a mi misma (?)
		if relacion.UsuarioRelacionID == ID {
			incluir = false
		}
		if incluir == true {
			usu.Password = ""
			usu.Biografia = ""
			usu.SitioWeb = ""
			usu.Ubicacion = ""
			usu.Banner = ""
			usu.Email = ""

			resultados = append(resultados, &usu)
		}

	}
	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return resultados, false
	}
	cursor.Close(contexto)
	return resultados, true
}
