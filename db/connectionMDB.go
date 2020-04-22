package db

import (
	"context"
	"log"

	"github.com/erikyvanov/API-Users-Posts/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnection es el objeto de conexión a la base de datos
var MongoConnection = ConnectMongo()

var clientOptions = options.Client().ApplyURI(config.URI)

// ConnectMongo es una funcion que conecta con la base de datos MongoDB
func ConnectMongo() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexión exitosa con la base de datos.")
	return client
}

// ConnectionOK dice si la conexión exite
func ConnectionOK() bool {
	err := MongoConnection.Ping(context.TODO(), nil)

	if err != nil {
		return false
	}

	return true
}
