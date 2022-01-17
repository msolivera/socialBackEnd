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
	router.HandleFunc("/modificarPerfil", middlew.ChequeoDB(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoDB(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweet", middlew.ChequeoDB(middlew.ValidoJWT(routers.LeoTweet))).Methods("GET")
	router.HandleFunc("/borrarTweet", middlew.ChequeoDB(middlew.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middlew.ChequeoDB(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/subirBanner", middlew.ChequeoDB(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoDB(middlew.ValidoJWT(routers.ObtenerAvatar))).Methods("GET")
	router.HandleFunc("/obtenerBanner", middlew.ChequeoDB(middlew.ValidoJWT(routers.ObtenerBanner))).Methods("GET")
	router.HandleFunc("/altaRelacion", middlew.ChequeoDB(middlew.ValidoJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.ChequeoDB(middlew.ValidoJWT(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlew.ChequeoDB(middlew.ValidoJWT(routers.ConsultaRelacion))).Methods("GET")
	router.HandleFunc("/listarUsuarios", middlew.ChequeoDB(middlew.ValidoJWT(routers.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/leerTweetSeguidores", middlew.ChequeoDB(middlew.ValidoJWT(routers.LeoTweetSeguidores))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
	//pongo a escuchar mi servidor
}
