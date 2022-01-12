package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/msolivera/socialTwittorBackEnd/bd"
)

func ObtenerAvatar(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	//tengo que traer los avatar de cualquier usuario
	if len(ID) < 1 {
		http.Error(w, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}
	//busco el perfil del id que recibi
	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "usuario no encontrado", http.StatusBadRequest)
		return
	}
	//aca intento abrir el archivo
	archivoMostrar, err := os.Open("uploads/avatars/" + perfil.Avatar)
	if err != nil {
		http.Error(w, "imagen no encontrada", http.StatusBadRequest)
		return
	}
	//copie el archivo
	_, err = io.Copy(w, archivoMostrar)
	if err != nil {
		http.Error(w, "error al copiar la imagen", http.StatusBadRequest)
	}
}
