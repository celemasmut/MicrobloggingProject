package jwt

import (
	"time"

	"github.com/celemasmut/MicrobloggingProject/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//GeneroJWT genera el encriptado con JWT
func GeneroJWT(usu models.Usuario) (string, error) {
	//Clave privada. puede ser cualquiera que solo nosotros lo sabemos
	miClave := []byte("SkillFactoryGo_Avalith")
	//se crea los claims
	payload := jwt.MapClaims{
		//nunca s epuede pasar el pass
		"email":             usu.Email,
		"nombre":            usu.Nombre,
		"apellidos":         usu.Apellidos,
		"fecha_naciemiento": usu.FechaNacimiento,
		"biografia":         usu.Biografia,
		"ubicacion":         usu.Ubicacion,
		"sitioweb":          usu.SitioWeb,
		"_id":               usu.ID.Hex(),
		"exp":               time.Now().Add(time.Hour * 24).Unix(), //formato unix es un formato rapido. tipo log
	}

	//obj de jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	//se realiza la firma con la clave privada
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
