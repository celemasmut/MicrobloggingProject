package routers

import (
	"encoding/json"
	"net/http"

	"github.com/celemasmut/MicrobloggingProject/bd"
	"github.com/celemasmut/MicrobloggingProject/models"
)

//Registro es la funcion para crear en la BD el registro de usuario
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t) // el body de http request es un stream. solo se lee una vez y luego se destruyee en memoria.
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar un pass de al menos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Mail ya registrado", 400)
		return
	}
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
