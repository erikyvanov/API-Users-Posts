package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User es el modelo de usuario que va a ser guardado en MongoDB
type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Avatar      string             `bson:"avatar" json:"avatar,omitempty"`
	Name        string             `bson:"name" json:"name,omitempty"`
	LastName    string             `bson:"lastname" json:"lastname,omitempty"`
	Description string             `bson:"description" json:"description,omitempty"`
	Birthday    time.Time          `bson:"birthday" json:"birthday,omitempty"`

	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password,omitempty"`
}
