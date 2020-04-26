package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ReturnPost es para enviar publicaciones al usuario
type ReturnPost struct {
	ID     primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID string             `bson:"userid" json:"userId,omitempty"`
	Body   string             `bson:"body" json:"body,omitempty"`
	Date   time.Time          `bson:"time" json:"time,omitempty"`
}
