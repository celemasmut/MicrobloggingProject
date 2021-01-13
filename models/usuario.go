package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Usuario es el model de usuario de la BD de mongoDB
type Usuario struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre          string             `bson:"_nombre" json:"nombre,omitempty"`
	Apellidos       string             `bson:"_apellidos" json:"apellidos,omitempty"`
	FechaNacimiento time.Time          `bson:"_fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email           string             `bson:"_email" json:"email"`
	Password        string             `bson:"_password" json:"password,omitempty"`
	Avatar          string             `bson:"_avatar" json:"avatar,omitempty"`
	Banner          string             `bson:"_banner" json:"banner,omitempty"`
	Biografia       string             `bson:"_biografia" json:"biografia,omitempty"`
	Ubicacion       string             `bson:"_ubicacion" json:"ubicacion,omitempty"`
	SitioWeb        string             `bson:"_sitioweb" json:"sitioweb,omitempty"`
}
