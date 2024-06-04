package db

// Ramtendo
//
// Francisco Alejandro Alcantar Aviles
// Marcos Ignacio Camacho Gonzalez
// Abraham Zumaya Manriquez
//
// package db
//
// Resto de las funciones CREATE, DELETE, UPDATE, GET de la base de datos

import (
	"context"

	pw "github.com/MarcosIgnacioo/playwright"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get de los términos de servicio
func GetTermsOfService() (bson.M, error) {
	var result bson.M
	coll := DB.Collection("terms_of_service")
	coll.FindOne(context.TODO(), bson.D{{"name", "service"}}).
		Decode(&result)
	return result, nil
}

// Get de los térmnos de privacidad
func GetTermsOfPrivacy() (bson.M, error) {
	var result bson.M
	coll := DB.Collection("terms_of_privacy")
	coll.FindOne(context.TODO(), bson.D{{"name", "privacy"}}).
		Decode(&result)
	return result, nil
}

// Insertar el kardex del alumno en la base de datos junto a su promedio general
func InsertKardex(name string, gpa int, kardex []interface{}) (*mongo.InsertOneResult, error) {
	_, checkStudent := GetKardex(name)
	if checkStudent != nil {
		coll := DB.Collection("kardex")
		r, e := coll.InsertOne(context.TODO(), pw.NewKardex(gpa, kardex, name))
		return r, e
	}
	return nil, nil
}

// Obtener el kardex del alumno
func GetKardex(name string) (bson.M, error) {
	var result bson.M
	coll := DB.Collection("kardex")
	err := coll.FindOne(context.TODO(), bson.D{{"name", name}}).
		Decode(&result)
	return result, err
}

// Insertar el mapa curricular del alumno
func InsertCurricularMap(name string, curricularMap []interface{}) (*mongo.InsertOneResult, error) {
	_, checkStudent := GetKardex(name)
	if checkStudent != nil {
		coll := DB.Collection("curricular_map")
		r, e := coll.InsertOne(context.TODO(), pw.NewCurricularMap(curricularMap, name))
		return r, e
	}
	return nil, nil
}

// Obtener el mapa curricular del alumno
func GetCurricularMap(name string) (bson.M, error) {
	var result bson.M
	coll := DB.Collection("curricular_map")
	err := coll.FindOne(context.TODO(), bson.D{{"name", name}}).
		Decode(&result)
	return result, err
}
