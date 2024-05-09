package db

import (
	"context"
	"github.com/MarcosIgnacioo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertProfile(post *models.Profile) (*mongo.InsertOneResult, error) {
	coll := DB.Collection("profiles")
	r, e := coll.InsertOne(context.TODO(), post)
	return r, e
}

func UpdateProfile(cn int, md *models.Profile) *mongo.SingleResult {
	coll := DB.Collection("profiles")
	update := bson.M{
		"$set": md,
	}
	err := coll.FindOneAndUpdate(context.TODO(), bson.D{{"control_number", cn}}, update)
	// _, err := coll.UpdateOne(context.TODO(), bson.D{{"control_number", cn}}, update)
	return err
}

func DeleteProfile(cn int) *mongo.SingleResult {
	coll := DB.Collection("profiles")
	// _, err := coll.DeleteOne(context.TODO(), bson.D{{"control_number", cn}})
	err := coll.FindOneAndDelete(context.TODO(), bson.D{{"control_number", cn}})
	return err
}

func GetProfile(cn int) (bson.M, error) {
	var result bson.M
	coll := DB.Collection("profiles")
	coll.FindOne(context.TODO(), bson.D{{"control_number", cn}}).
		Decode(&result)
	return result, nil
}

func LikePost() {
}
