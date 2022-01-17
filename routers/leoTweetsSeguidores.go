package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/msolivera/socialTwittorBackEnd/bd"
)

func LeoTweetSeguidores(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "deve enviar numero", http.StatusBadRequest)
		return
	}
	respuesta, correcto := bd.LeoTweetsSeguidores(IDUsuario, pagina)
	if correcto == false {
		http.Error(w, "error al leer los tweet", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
