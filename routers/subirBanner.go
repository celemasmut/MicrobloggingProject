package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/celemasmut/MicrobloggingProject/bd"
	"github.com/celemasmut/MicrobloggingProject/models"
)

//SubirBanner sube el avatar al servidor
func SubirBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banners")

	var extension = strings.Split(handler.Filename, ".")[1]

	var archivo string = "uploads/banners/" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen !"+err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(f, file) //se graba en disco
	if err != nil {
		http.Error(w, "Error al copiar la imagen !"+err.Error(), http.StatusBadRequest)
		return
	}

	var usu models.Usuario
	var status bool
	usu.Banner = IDUsuario + "." + extension
	status, err = bd.ModificoRegistro(usu, IDUsuario)
	if err != nil || status == false {
		http.Error(w, "Error al grabar el Banner en la BD !"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
