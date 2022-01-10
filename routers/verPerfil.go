package routers

import (
	"encoding/json"
	"net/http"

	"github.com/msolivera/socialTwittorBackEnd/bd"
)

//permite extraer los valores del perfil
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	//no devuelve nada
	//extraemos del body del request los datos, los extraemos del url
	ID := r.URL.Query().Get("id")
	//si no encuentra el id
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el ID", http.StatusBadRequest)
		return
	}
	//si anduvo bien creo perfil que por la funcion de buscoPerfil es un modelo
	perfil, err := bd.BuscoPerfil(ID)
	//si no encuentro el perfil del usuario
	if err != nil {
		http.Error(w, "ocurrio un error al buscar el registro "+err.Error(), 400)
		return
	}
	//si funciona todo seteamos el header y el conext y devuelvo el json
	w.Header().Set("context-type", "application/json")
	//devolvemos el status de created (todo ok)
	w.WriteHeader(http.StatusCreated)
	//retorna el json de perfil
	json.NewEncoder(w).Encode(perfil)

}
