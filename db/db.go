package db

// Ramtendo
//
// Francisco Alejandro Alcantar Aviles
// Marcos Ignacio Camacho Gonzalez
// Abraham Zumaya Manriquez
//
// package db
// Aqui es donde se establece conexion con la base de datos

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CREDENTIALS      = "credentials"
	KARDEX           = "kardex"
	CURRICULAR_MAP   = "curricular_map"
	MOODLE           = "moodle"
	CLASSROOM        = "classroom"
	ACTIVITIES       = "activities"
	TERMS_OF_SERVICE = "terms_of_service"
	TERMS_OF_PRIVACY = "terms_of_privacy"
	CALENDAR         = "calendar"
)

var client *mongo.Client
var clientErr error
var DB *mongo.Database

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("DB_NAME")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " +
			"www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, clientErr = mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if clientErr != nil {
		panic(clientErr)
	}
	DB = client.Database(dbName)
}
