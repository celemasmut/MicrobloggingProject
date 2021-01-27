package bd

import (
	"context"
	"log"
	"time"

	"github.com/celemasmut/MicrobloggingProject/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//LeoTweets es una funcion que lee los tweets de un perfil.
func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microdb")
	col := db.Collection("tweet")

	var resultados []*models.DevuelvoTweets

	condicion := bson.M{
		"userid": ID, //userid es el id de mongo. el ID ya como string es el que se pasa desde routers a esta rutina
	}

	opciones := options.Find()
	opciones.SetLimit(20)                               //indica cuantos doc va a traer por limite que va por parametro
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}}) //me trae de forma descendente. donde los ultimos seran los primeros en mostrar
	opciones.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones) //cursor es un obj json
	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	for cursor.Next(context.TODO()) {
		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}
	return resultados, true
}
