package models

// SendPost es una estructura para mandar un post en json
type SendPost struct {
	Message string `bson:"messaje" json:"message"`
}
