package routers

import (
	"ecoding/json"
	"encoding/json"
	"net/http"

	"github.com/msolivera/socialTwittorBackEnd/bd"
	"github.com/msolivera/socialTwittorBackEnd/models"
)

//para crear registro de usuario en BD
func Registro(w http.ResponseWriter, r *http.Request) {

	var usu models.Usuario
	//.Body es un objeto de tipo Stream, despues de que se lee se destruye
	err := json.NewDecoder(r.Body).Decode(&usu)

	if err != nil {
		//borrar esto despues esta mal
		bd.ChequeoConexion()

	}

}
