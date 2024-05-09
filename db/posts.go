package db

import (
	"context"
	"github.com/MarcosIgnacioo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertPost(post *models.Post) (*mongo.InsertOneResult, error) {
	coll := DB.Collection("posts")
	r, e := coll.InsertOne(context.TODO(), post)
	return r, e
}

func UpdatePost(cn int, md *models.Post) *mongo.SingleResult {
	coll := DB.Collection("posts")
	update := bson.M{
		"$set": md,
	}
	err := coll.FindOneAndUpdate(context.TODO(), bson.D{{"control_number", cn}}, update)
	// _, err := coll.UpdateOne(context.TODO(), bson.D{{"control_number", cn}}, update)
	return err
}

func DeletePost(cn int) *mongo.SingleResult {
	coll := DB.Collection("posts")
	// _, err := coll.DeleteOne(context.TODO(), bson.D{{"control_number", cn}})
	err := coll.FindOneAndDelete(context.TODO(), bson.D{{"control_number", cn}})
	return err
}

func GetPost(cn int) (bson.M, error) {
	var result bson.M
	coll := DB.Collection("posts")
	coll.FindOne(context.TODO(), bson.D{{"control_number", cn}}).
		Decode(&result)
	return result, nil
}
