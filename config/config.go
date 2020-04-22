package config

/*Es seccion es para tener control de las costantes que se usan en todo el codigo*/

// URI es donde esta alojada la base de datos
const URI string = "mongodb://localhost:27017"

// DBName es el nombre de la base de datos
const DBName string = "API-Users-Posts"

// TokenKey es la clave para usar uso de los tokens
var TokenKey = []byte("HolaGithub")
