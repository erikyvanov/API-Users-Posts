package db

import (
	"context"
	"fmt"
	"time"

	"github.com/erikyvanov/API-Users-Posts/config"
	"github.com/erikyvanov/API-Users-Posts/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReadUsers lee todos los uaurios en la base de datos
// Recibe el ID del usuario, la pagina, el nombre de quien va a buscar, el tipo de busqueda
func ReadUsers(ID string, page int64, search string, tipo string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dataBase := MongoConnection.Database(config.DBName)
	collection := dataBase.Collection("users")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := collection.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var found, include bool

	for cur.Next(ctx) {
		var s models.User
		err = cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error() + "Despues de hacer el decode")
			return results, false
		}

		var r models.Relation
		r.UserID = ID
		r.RelationID = s.ID.Hex()

		include = false
		found, _ = ConsultRelation(r)

		// Si no lo sigo y lo estoy buscando
		if tipo == "new" && !found {
			include = true
		}
		// Si solo quiero buscar a los usuarios que sigo
		if tipo == "follow" && found {
			include = true
		}

		if r.RelationID == ID {
			include = false
		}

		// Si lo debo incluir
		if include {
			s.Password = ""
			s.Email = ""

			s.Description = ""

			results = append(results, &s)
		}
	}

	if cur.Err() != nil {
		fmt.Println(err.Error())
		return results, false
	}

	cur.Close(ctx)
	return results, true
}
