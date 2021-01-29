package bd

import (
	"context"
	"time"

	"github.com/celemasmut/MicrobloggingProject/models"
)

//InsertoRelacion graba la relacion en la BD
func InsertoRelacion(usu models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("microdb")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, usu)
	if err != nil {
		return false, err
	}
	return true, nil
}
