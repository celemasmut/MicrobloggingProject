package middlew

import (
	"net/http"

	"github.com/celemasmut/MicrobloggingProject/bd"
)

//ChequeoBD es el middleware que me permite conocer el estado de la BD
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc { //handlerFunc son manejadores de funciones
	return func(w http.ResponseWriter, r *http.Request) { //retorno de funcion anonima
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion Perdida con la BD", 500) //recibe un status 500  diciendo que el error es de parte del servidor y no continua
			return
		}
		next.ServeHTTP(w, r) //si no da error pasa todos los valores al prox eslabon de la cadena
	}
}
