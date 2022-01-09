package jsonwebtoken

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/msolivera/socialTwittorBackEnd/models"
)

func GeneroToken(usu models.Usuario) (string, error) {
	miClave := []byte("miClaveParaEstaParte")

	//cosas que van en el cuerpo del json
	payload := jwt.MapClaims{
		"email":           usu.Email,
		"nombre":          usu.Nombre,
		"apellidos":       usu.Apellidos,
		"fechaNacimiento": usu.FechaNacimiento,
		"biografia":       usu.Biografia,
		"ubicacion":       usu.Ubicacion,
		"sitioWeb":        usu.SitioWeb,
		"_id":             usu.ID.Hex(),
		"exp":             time.Now().Add(time.Hour * 24).Unix(),
	}

	//funcion de jwt, le paso el algoritmo del token y el cuerpo
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	//lo firmo con mi clave
	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
