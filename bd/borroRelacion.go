package bd

import (
	"context"
	"time"

	"github.com/celemasmut/MicrobloggingProject/models"
)

//BorroRelacion elimina la relacion de la BD
func BorroRelacion(usu models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microdb")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, usu)
	if err != nil {
		return false, err
	}
	return true, nil
}
