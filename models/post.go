package models

import "time"

// Post es la estructura de una publicacion en la app
type Post struct {
	UserID string    `bson:"userid" json:"userid, omitempty"`
	Body   string    `bson:"body" json:"body, omitempty"`
	Date   time.Time `bson:"date" json:"date, omitempty"`
}
