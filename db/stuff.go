package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
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
