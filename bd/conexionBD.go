package bd

import (
	"context" // siven para comunicar info entre ejecucion y ejecucion y ademas permite setear serie de valores como  un timeOut
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN es el objeto de conexion a la BD
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://microuser:rcm4741695@microblog.lemq3.mongodb.net/microGo?retryWrites=true&w=majority")

//ConectarBD permite conectar la BD
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil) //llama a la BD para saber si la BD esta arriba.
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa con la BD")
	return client
}

//ChequeoConnection es el ping a la BD
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
