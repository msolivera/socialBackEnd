package routers

import (
	"encoding/json"
	"net/http"

	"github.com/msolivera/socialTwittorBackEnd/bd"
	"github.com/msolivera/socialTwittorBackEnd/models"
)

func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	t := models.Relacion{
		UsuarioID:         IDUsuario,
		UsuarioRelacionID: ID,
	}

	var respuesta models.RespuestaConsultaRelacion

	status, err := bd.ConsultoRelacion(t)

	if err != nil || status == false {
		respuesta.Status = false
	} else {
		respuesta.Status = true
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
