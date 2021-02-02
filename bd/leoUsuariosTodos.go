package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/celemasmut/MicrobloggingProject/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//LeoUsuariosTodos lee los usuarios registrados en el sistma, si se ecibe "R" en quienes
//trae solos los que se relacionan conmigo
func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microdb")
	col := db.Collection("usuarios")

	var results []*models.Usuario

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var encontrado, incluir bool
	for cur.Next(ctx) {
		var s models.Usuario
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var relacion models.Relacion
		relacion.UsuarioID = ID
		relacion.UsuarioRelacionID = s.ID.Hex()

		incluir = false
		encontrado, err = ConsultoRelacion(relacion)
		if tipo == "new" && encontrado == false {
			incluir = true
		}
		if tipo == "follow" && encontrado == true {
			incluir = true
		}
		if relacion.UsuarioRelacionID == ID {
			incluir = false
		}
		if incluir == true {
			s.Password = ""
			s.Biografia = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cur.Close(ctx)

	return results, true

}
