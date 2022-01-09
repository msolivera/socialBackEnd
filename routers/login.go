package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/msolivera/socialTwittorBackEnd/bd"
	jsonwebtoken "github.com/msolivera/socialTwittorBackEnd/jsonWebToken"
	"github.com/msolivera/socialTwittorBackEnd/models"
)

//realizacion del login que recibe el endpoint
func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")

	var usu models.Usuario
	err := json.NewDecoder(r.Body).Decode(&usu)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña invalidos"+err.Error(), 400)
		return
	}
	if len(usu.Email) == 0 {
		http.Error(w, "Email es campo requerido", 400)
		return
	}

	documento, existe := bd.IntentoLogin(usu.Email, usu.Password)
	if existe == false {
		http.Error(w, "Usuario y/o contraseña invalidos", 400)
		return
	}

	//genero una variable que va a tener un json web token
	jwtkey, err := jsonwebtoken.GeneroToken(documento)

	//si ocurre un error lo deuelvo
	if err != nil {
		http.Error(w, "Ocurrio un eror "+err.Error(), 400)
	}
	//sino, genero una respuesta que va a asignar a un json el nuevo token
	resp := models.RespuestaLogin{
		Token: jwtkey,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	//esto no se va a usar pero vamso a ver como grabar una cookie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{Name: "Token", Value: jwtkey, Expires: expirationTime})

}
