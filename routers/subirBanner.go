package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/msolivera/socialTwittorBackEnd/bd"
	"github.com/msolivera/socialTwittorBackEnd/models"
)

func SubirBanner(w http.ResponseWriter, r *http.Request) {
	//capturo archivo desde el request
	file, handler, err := r.FormFile("banner")

	//capturo la extension del archivo
	var extension = strings.Split(handler.Filename, ".")[1]
	//formo la ruta y el nombre donde voy a grabar el archivo
	var archivo string = "uploads/banners/" + IDUsuario + "." + extension
	//funcion del ssoo,abro el archivo y le doy permiso de escritura, modificacion y lectura
	//creando en el disco
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "error al subir el Banner "+err.Error(), http.StatusBadRequest)
		return
	}
	//luego de reservar el espacio en disco para subir el archivo tengo que copiarlo
	//toma file que viene del r y lo copia donde creamos el espacio en disco
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "error al copiar el Banner "+err.Error(), http.StatusBadRequest)
	}
	//modifico el registro del usuario
	var usuario models.Usuario
	var status bool
	//le asigno al usuario la ruta de la imagen
	usuario.Banner = IDUsuario + "." + extension
	//modifico el registro con el nuevo dato
	status, err = bd.ModificoRegistro(usuario, IDUsuario)
	if err != nil || status == false {
		http.Error(w, "error al grabar en la BD "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
