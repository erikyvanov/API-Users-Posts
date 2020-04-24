# API-Users-Posts
Es una API para servir a una app con sistema de usuarios y publicaciones.

## Requiere de:
### Mongo Driver
- go get go.mongodb.org/mongo-driver/mongo
### Crypto Bcrypt
- go get golang.org/x/crypto/bcrypt
### Gorilla Mux
- go get github.com/gorilla/mux
### RS Cors
- go get github.com/rs/cors
### Dgrijalva JWT-GO
- go get github.com/dgrijalva/jwt-go

## Configuración:
Existe la configuaración de la base de datos y de las claves para los JWT.
Hay que irse a la carpeta config.
### Base de datos:
Por defecto se usa la base de datos local predetermianda de MongoDB y el nombre de la base de datos es "API-Users-Posts".
```go
// URI es donde esta alojada la base de datos
const URI string = "mongodb://localhost:27017"

// DBName es el nombre de la base de datos
const DBName string = "API-Users-Posts"
```
### JWT
Por defecto la clave de los tokens es "HolaGithub". Se recomienda cambiarla para mayor seguridad.
```go
// TokenKey es la clave para usar uso de los tokens
var TokenKey = []byte("HolaGithub")
```
