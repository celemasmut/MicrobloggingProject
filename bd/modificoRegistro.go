package bd

import (
	"context"
	"time"

	"github.com/celemasmut/MicrobloggingProject/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ModificoRegistro permite modificar el perfil del usuario
func ModificoRegistro(usu models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microdb")
	col := db.Collection("usuarios")

	registro := make(map[string]interface{})

	if len(usu.Nombre) > 0 {
		registro["nombre"] = usu.Nombre
	}
	if len(usu.Apellidos) > 0 {
		registro["apellidos"] = usu.Apellidos
	}

	registro["fechaNacimiento"] = usu.FechaNacimiento

	if len(usu.Avatar) > 0 {
		registro["avatar"] = usu.Avatar
	}
	if len(usu.Banner) > 0 {
		registro["banner"] = usu.Banner
	}
	if len(usu.Biografia) > 0 {
		registro["biografia"] = usu.Biografia
	}
	if len(usu.Ubicacion) > 0 {
		registro["ubicacion"] = usu.Ubicacion
	}
	if len(usu.SitioWeb) > 0 {
		registro["sitioweb"] = usu.SitioWeb
	}

	updtString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)

	if err != nil {
		return false, err
	}

	return true, nil
}
