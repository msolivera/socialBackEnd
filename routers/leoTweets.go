package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/msolivera/socialTwittorBackEnd/bd"
)

func LeoTweet(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "debe enviar el ID", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "debe enviar el numero de pagina", http.StatusBadRequest)
		return
	}

	//convertir letra a int
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "pagina debe ser mayor a cero", http.StatusBadRequest)
		return
	}
	pag := int64(pagina)

	respuesta, correcto := bd.LeoTweet(ID, pag)
	if correcto == false {
		http.Error(w, "ocurrio un error", http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

}
