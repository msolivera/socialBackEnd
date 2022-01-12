package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/msolivera/socialTwittorBackEnd/bd"
)

func ObtenerBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	//tengo que traer los avatar de cualquier usuario
	if len(ID) < 1 {
		http.Error(w, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}
	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "usuario no encontrado", http.StatusBadRequest)
		return
	}

	archivoMostrar, err := os.Open("uploads/banners/" + perfil.Banner)
	if err != nil {
		http.Error(w, "imagen no encontrada", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, archivoMostrar)
	if err != nil {
		http.Error(w, "error al copiar la imagen", http.StatusBadRequest)
	}
}
