package bd

import (
	"context"
	"time"

	"github.com/celemasmut/MicrobloggingProject/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertoTweet graba el Tweet en la BD
func InsertoTweet(usu models.GraboTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microdb")
	col := db.Collection("tweet")

	registro := bson.M{
		"userid":  usu.UserID,
		"mensaje": usu.Mensaje,
		"fecha":   usu.Fecha,
	}

	result, err := col.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
