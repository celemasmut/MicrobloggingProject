package bd

import (
	"context"
	"time"

	"github.com/celemasmut/MicrobloggingProject/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertoRegistro es la parada final con la BD para insertar los datos del usuario
func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microdb")
	col := db.Collection("usuarios")
	//encriptamos la pass
	u.Password, _ = EncriptarPassword(u.Password)
	//col es de tipo coleccion
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	//InsertedID es una funcion que devuelve un objeto de tipo ObjectID
	//objID es un objeto de tipo ObjectID
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil //como retorna un string se convierte el obj a un string
}
