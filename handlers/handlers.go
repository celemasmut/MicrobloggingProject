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
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoBD(routers.ObtenerAvatar)).Methods("GET") //ObtenerAvatar/Banner no necesitan la validacion de Token

	router.HandleFunc("/subirBanner", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.ChequeoBD(routers.ObtenerBanner)).Methods("GET")

	router.HandleFunc("/altaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.BajaRelacion))).Methods("DELETE")

	router.HandleFunc("/consultaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.ConsultaRelacion))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
