package models

// Relation es un a estructura para hacer relaciones entre los usuarios
type Relation struct {
	UserID     string `bson:"userid" json:"userID"`
	RelationID string `bson:"relationid" json:"relationId"`
}
