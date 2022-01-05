package middlew

import (
	"net/http"

	"github.com/msolivera/socialTwittorBackEnd/bd"
)

//Un middleware tiene que recibir algo y devolver lo mismo, no puede devolver tipos distintos
func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConexion() == 0 {
			http.Error(w, "Conexion perdida con la BD", 500)
			return
		}
		//si dio error me manda un return
		//si no da error paso al proximo eslabon de la cadena la informacion
		next.ServeHTTP(w, r)
	}

}
