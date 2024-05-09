package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/MarcosIgnacioo/models"
	pw "github.com/MarcosIgnacioo/playwright"
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

func Mongo() {
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// coll := DB.Collection("students")

	ramsesupdate := pw.NewStudentInfo(999, "ramsesactualizado", "ramses@gmail.com", "La paz", "2024-I", "Ing Software", 6, "A", "M", "Active")

	// r, e := coll.InsertOne(context.TODO(), ramses)

	// class_subject	"Ingeniería de Software II Turno Matutino"
	// title	"Trabajo de manteniemiento está en fecha de entrega"
	// link	"https://enlinea2024-1.uabcs.mx/mod/assign/view.php?id=85998"
	// date
	// day	"2"
	// month	"junio"
	// year	"2024"
	// hour	"00:00"

	before, _ := GetStudent(1)
	fmt.Println(before)
	UpdateStudent(1123123213, ramsesupdate)
	DeleteStudent(999)
	// after, _ := GetStudent(999)
	// fmt.Println(after)

	// err := coll.FindOne(context.TODO(), bson.D{{"name", name}}).
	// 	Decode(&result)
	// if err == mongo.ErrNoDocuments {
	// 	fmt.Printf("No document was found with the name %s\n", name)
	// 	return
	// }
	// if err != nil {
	// 	panic(err)
	// }
	// jsonData, err := json.MarshalIndent(result, "", "    ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", jsonData)
}
