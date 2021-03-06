package middlew

import (
	"net/http"

	"github.com/celemasmut/MicrobloggingProject/routers"
)

// este seria el Middleware que va a chequear la validez del JWT,

//ValidoJWT permite validar el JWT que nos viene en la peticion
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token !"+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}

//este middlew se usara en Handlers
