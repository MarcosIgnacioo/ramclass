package db

import (
	"context"
	"github.com/MarcosIgnacioo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var MOODLE_COLLECTION = "moodle"

func InsertMoodle(student *models.Moodle) (*mongo.InsertOneResult, error) {
	coll := DB.Collection(MOODLE_COLLECTION)
	r, e := coll.InsertOne(context.TODO(), student)
	return r, e
}

func UpdateMoodle(cn int, md *models.Moodle) error {
	coll := DB.Collection(MOODLE_COLLECTION)
	update := bson.M{
		"$set": md,
	}
	_, err := coll.UpdateOne(context.TODO(), bson.D{{"control_number", cn}}, update)
	return err
}

func DeleteMoodle(cn int) error {
	coll := DB.Collection(MOODLE_COLLECTION)
	_, err := coll.DeleteOne(context.TODO(), bson.D{{"control_number", cn}})
	return err
}

func GetMoodle(cn int) (bson.M, error) {
	var result bson.M
	coll := DB.Collection(MOODLE_COLLECTION)
	coll.FindOne(context.TODO(), bson.D{{"control_number", cn}}).
		Decode(&result)
	return result, nil
}
