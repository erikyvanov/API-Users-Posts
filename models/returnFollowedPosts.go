package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ReturnFollowePost regresa un post de un usuario
type ReturnFollowePost struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID     string             `bson:"userid" json:"userId,omitempty"`
	RelationID string             `bson:"relationid" json:"relationId,omitempty"`

	Posts struct {
		Body string    `bson:"body" json:"body,omitempty"`
		Date time.Time `bson:"date" json:"date,omitempty"`
		ID   string    `bson:"_id" json:"_id,omitempty"`
	}
}
