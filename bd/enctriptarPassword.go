package bd

import "golang.org/x/crypto/bcrypt"

func EncriptarPassword(pass string) (string, error) {

	//costo significa que el algoritmo va a pasar 2^costo veces sobre
	//la pass que pasemos para encriptar
	//a mayor costo mas segura
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
