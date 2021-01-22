package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/celemasmut/MicrobloggingProject/bd"
	"github.com/celemasmut/MicrobloggingProject/models"

	jwt "github.com/celemasmut/MicrobloggingProject/jwt"
)

//Login realiza el login de usuario
func Login(w http.ResponseWriter, r *http.Request) {
	//se setea en el header que en contenido  que se devuelva (w)
	//se tipo json
	w.Header().Add("content-type", "application/json")

	var usu models.Usuario

	err := json.NewDecoder(r.Body).Decode(&usu)
	if err != nil {
		http.Error(w, "Usuario y/o Contraseña invalidos"+err.Error(), 400)
		return // cancela el endpoint y o hace mas nada
	}

	if len(usu.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	documento, existe := bd.IntentoLogin(usu.Email, usu.Password)
	if existe == false {
		http.Error(w, "Usuario y/o Contraseña invalidos", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el Token correspondiente"+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	//como grabar una cookie
	//se genera campo fecha para ver la expiracion de la cookie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
