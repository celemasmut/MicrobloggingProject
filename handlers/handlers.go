package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/celemasmut/MicrobloggingProject/middlew"
	"github.com/celemasmut/MicrobloggingProject/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Manejadores seteo el puerto en Handler y pongo a escuchar el servidor
func Manejadores() {
	router := mux.NewRouter()

	//primer parametro es donde se va a dirigir, el middlew chequea con  POST
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
