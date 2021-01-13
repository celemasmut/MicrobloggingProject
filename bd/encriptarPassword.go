package bd

import "golang.org/x/crypto/bcrypt"

//EncriptarPassword es la funcion que me permite encriptar la password
func EncriptarPassword(pass string) (string, error) {
	costo := 8 //algoritmo basado en dos elevado al costo. No bajar del 6 lo aconsejable es 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
