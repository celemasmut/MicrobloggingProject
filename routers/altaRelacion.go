package routers

import (
	"net/http"

	"github.com/celemasmut/MicrobloggingProject/bd"
	"github.com/celemasmut/MicrobloggingProject/models"
)

//AltaRelacion realiza el registro de la relacion entre usuarios
func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "el parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var usu models.Relacion

	usu.UsuarioID = IDUsuario
	usu.UsuarioRelacionID = ID
	status, err := bd.InsertoRelacion(usu)
	if err != nil {
		http.Error(w, "Ocurrio un error al intenatr insertar relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado intenatr insertar relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
