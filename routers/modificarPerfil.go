package routers

import (
	"encoding/json"
	"net/http"

	"github.com/celemasmut/MicrobloggingProject/bd"
	"github.com/celemasmut/MicrobloggingProject/models"
)

//ModificarPerfil modifica el perfil del usuario
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var usu models.Usuario

	err := json.NewDecoder(r.Body).Decode(&usu)

	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.ModificoRegistro(usu, IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modficar el registro, reintente nuevamente"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado modficar el registro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
