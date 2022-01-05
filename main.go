package main

import (
	"log"

	"github.com/msolivera/socialTwittorBackEnd/bd"
)

func main() {
	if bd.ChequeoConexion() == 0 {
		log.Fatal("Sin conexion")
	}

}
