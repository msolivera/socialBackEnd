package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/msolivera/socialTwittorBackEnd/bd"
	"github.com/msolivera/socialTwittorBackEnd/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {

	var mensaje models.Tweet

	//recibo el body, lo decodificamos y lo grabamos en mensaje
	err := json.NewDecoder(r.Body).Decode(&mensaje)
	registro := models.GraboTweet{
		UsuarioID: IDUsuario,
		Mensaje:   mensaje.Mensaje,
		Fecha:     time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "ocurrio un error al intentar ingresar el regustro "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se logro insertar el tweet ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
