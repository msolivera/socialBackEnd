package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/msolivera/socialTwittorBackEnd/bd"
)

func ListaUsuarios(w http.ResponseWriter, r *http.Request) {
	tipoUsu := r.URL.Query().Get("tipo")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}
	pag := int64(pageTemp)
	result, status := bd.LeoUsuariosTodos(IDUsuario, pag, search, tipoUsu)
	if status == false {
		http.Error(w, "Error al leer a los usuarios", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}
