package routers

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/msolivera/socialTwittorBackEnd/models"
	"github.com/msolivera/socialTwittorBackEnd/bd"
	"github.com/msolivera/socialTwittorBackEnd/jwt"
)

//realizacion del login que recibe el endpoint
func Login(w http.ResponseWriter, r *http.Request){

	w.Header().Add("content-type","application/json")

	var usu models.Usuario
	err := json.NewDecoder(r.Body.Decode(&usu)
	if err != nil{
		http.Error(w, "Usuario y/o contraseña invalidos"+err.Error(),400)
		return
	}
	if len(usu.Email)==0{
		http.Error(w, "Email es campo requerido" ,400)
		return
	}

	/*documento, existe := bd.
	
	bd.IntentoLogin(usu.Email, usu.Password){
		if existe == false{
			http.Error(w, "Usuario y/o contraseña invalidos"+err.Error(),400)
		return
		}
	}*/
}