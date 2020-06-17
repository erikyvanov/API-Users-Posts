package db

import (
	"context"
	"time"

	"github.com/erikyvanov/API-Users-Posts/config"
	"github.com/erikyvanov/API-Users-Posts/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ModifyRegister modifica el perfil de un usuario en la base de datos
//Regresa con un bool si se hicieron los cambios o un error en caso de que no
func ModifyRegister(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dataBase := MongoConnection.Database(config.DBName)
	collection := dataBase.Collection("users")

	// Mapear los cambios
	updates := make(map[string]interface{})
	if len(u.Avatar) > 0 {
		updates["avatar"] = u.Avatar
	}
	if len(u.LastName) > 0 {
		updates["lastname"] = u.LastName
	}
	if len(u.Name) > 0 {
		updates["name"] = u.Name
	}
	updates["description"] = u.Description
	updates["birthday"] = u.Birthday
	updates["location"] = u.Location

	// Meter los cambios que se van a hacer
	updtString := bson.M{
		"$set": updates,
	}
	// ID a tipo bson.ObjectID
	objID, _ := primitive.ObjectIDFromHex(ID)

	//El filtro para encontrar el usuario, con un $eq para que sea equivalente al ID
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	// Se hace la actualizacion de datos
	_, err := collection.UpdateOne(ctx, filter, updtString)
	if err != nil {
		return false, err
	}

	return true, nil
}
