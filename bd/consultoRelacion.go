package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/celemasmut/MicrobloggingProject/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ConsultoRelacion consulta la relacion entre 2 usuarios
func ConsultoRelacion(usu models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("microdb")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         usu.UsuarioID,
		"usuariorelacionid": usu.UsuarioRelacionID,
	}

	var resultado models.Relacion
	fmt.Println(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil

}
