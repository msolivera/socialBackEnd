package routers

import (
	"encoding/json"
	"net/http"

	"github.com/msolivera/socialTwittorBackEnd/bd"
	"github.com/msolivera/socialTwittorBackEnd/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var usu models.Usuario

	err := json.NewDecoder(r.Body).Decode(&usu)
	if err != nil {
		http.Error(w, "datos incorrectos "+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.ModificoRegistro(usu, IDUsuario)

	if err != nil {
		http.Error(w, "ocurrio un error al intentar modificar registro "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "no se ha logrado modificar el registro "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
