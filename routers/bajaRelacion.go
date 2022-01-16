package routers

import (
	"net/http"

	"github.com/msolivera/socialTwittorBackEnd/bd"
	"github.com/msolivera/socialTwittorBackEnd/models"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	t := models.Relacion{
		UsuarioID:         IDUsuario,
		UsuarioRelacionID: ID,
	}

	status, err := bd.BorroRelacion(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al borrar relacion "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se logro borrar relacion "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
