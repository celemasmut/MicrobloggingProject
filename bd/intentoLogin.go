package bd

import (
	"github.com/celemasmut/MicrobloggingProject/models"
	"golang.org/x/crypto/bcrypt"
)

//IntentoLogin realiza el chequeo de login a la BD
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	//verifico si existe
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usu, false
	}
	//grabo en dos variables la pass que vino por parametro y la que esta grabada en la BD
	//dos slice de byte
	passwordBytes := []byte(password)  //esta no va a estar encriptada
	passwordBD := []byte(usu.Password) //va a estar encriptada

	//con esta funcion verificamos si la pass es correcta.
	//primero va la pass encritada de la BD
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return usu, false
	}
	return usu, true
}
