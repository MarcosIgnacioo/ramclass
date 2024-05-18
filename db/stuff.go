package db

import (
	"context"

	pw "github.com/MarcosIgnacioo/playwright"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get del calendario escolar

// Get de los terminos de servicio y privacidad

func GetTermsOfService() (bson.M, error) {
	var result bson.M
	coll := DB.Collection("terms_of_service")
	coll.FindOne(context.TODO(), bson.D{{"name", "service"}}).
		Decode(&result)
	return result, nil
}

func GetTermsOfPrivacy() (bson.M, error) {
	var result bson.M
	coll := DB.Collection("terms_of_privacy")
	coll.FindOne(context.TODO(), bson.D{{"name", "privacy"}}).
		Decode(&result)
	return result, nil
}

func GetSchoolCalendar() (bson.M, error) {
	var result bson.M
	coll := DB.Collection("calendar")
	coll.FindOne(context.TODO(), bson.D{{"name", "calendar"}}).
		Decode(&result)
	return result, nil
}

func InsertKardex(name string, gpa int, kardex []interface{}) (*mongo.InsertOneResult, error) {
	_, checkStudent := GetKardex(name)
	if checkStudent != nil {
		coll := DB.Collection("kardex")
		r, e := coll.InsertOne(context.TODO(), pw.NewKardex(gpa, kardex, name))
		return r, e
	}
	return nil, nil
}

func GetKardex(name string) (bson.M, error) {
	var result bson.M
	coll := DB.Collection("kardex")
	err := coll.FindOne(context.TODO(), bson.D{{"name", name}}).
		Decode(&result)
	return result, err
}

func InsertCurricularMap(name string, curricularMap []interface{}) (*mongo.InsertOneResult, error) {
	_, checkStudent := GetKardex(name)
	if checkStudent != nil {
		coll := DB.Collection("curricular_map")
		r, e := coll.InsertOne(context.TODO(), pw.NewCurricularMap(curricularMap, name))
		return r, e
	}
	return nil, nil
}

func GetCurricularMap(name string) (bson.M, error) {
	var result bson.M
	coll := DB.Collection("curricular_map")
	err := coll.FindOne(context.TODO(), bson.D{{"name", name}}).
		Decode(&result)
	return result, err
}
