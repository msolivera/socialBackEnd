package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/msolivera/socialTwittorBackEnd/middlew"
	"github.com/msolivera/socialTwittorBackEnd/routers"
	"github.com/rs/cors"
)

/*cuando ejecutemos la API va a llamar esta funcion
ACA VOY A VER LAS RUTAS DE LA API
*/
func Manejadores() {

	router := mux.NewRouter()
	//Asi es como se definen las rutas.
	//el HandrleFunc sabe que si se pone determinada ruta va a llamar al middleware que hace determinanda funcion
	//en este caso el middleware chequea la BD, si esta OK, se continua con la ejecucion
	//y luego hay que ponerle que tipo de metodo va a ser el que se pida en esa ruta.
	router.HandleFunc("/registro", middlew.ChequeoDB(routers.Registro)).Methods("POST")
	//comparten misma cadena de middleware
	router.HandleFunc("/login", middlew.ChequeoDB(routers.Login)).Methods("POST")
	///////
	router.HandleFunc("/verPerfil", middlew.ChequeoDB(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
	//pongo a escuchar mi servidor
}
