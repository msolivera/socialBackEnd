package routers

import (
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
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return

	}
	//verifico si vino el email en el usuario
	if len(usu.Email) == 0 {
		http.Error(w, "El email es requerido", 400)
		return
	}
	if len(usu.Password) < 6 {
		http.Error(w, "La contraseña debe tener al menos 6 caracteres", 400)
		return
	}

	//esto para verificar que el email no este repetido
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(usu.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ene email", 400)
		return
	}

	//esto es para cuando no sabemos que error salio
	_, status, err := bd.InsertoRegistro(usu)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}

	//esta es una ultima confirmacion por si ocurre algo al insertar los datos
	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
