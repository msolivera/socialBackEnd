package bd

import (
	"github.com/msolivera/socialTwittorBackEnd/models"
	"golang.org/x/crypto/bcrypt"
)

//realiza el chequeo de login a la BD
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usu, false
	}
	//conversion en bytes de lo que recibo por parametro
	passwordBytes := []byte(password)
	//password que trae el modelo de usuario de la BD
	passwordBD := []byte(usu.Password)

	//comparo ambas contraseñas
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
