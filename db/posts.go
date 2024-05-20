package db

import (
	"context"

	"github.com/MarcosIgnacioo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertTasks(tasks *models.Tasks) error {
	coll := DB.Collection("tasks")
	res := coll.FindOneAndReplace(context.TODO(), bson.D{{"identifier", tasks.Identifier}}, tasks)
	if res.Err() != nil {
		coll.InsertOne(context.TODO(), tasks)
	}
	return res.Err()
}

func GetTasks(identifier string) (bson.M, error) {
	var result bson.M
	coll := DB.Collection("tasks")
	coll.FindOne(context.TODO(), bson.D{{"identifier", identifier}}).
		Decode(&result)
	return result, nil
}

func UpdateTask(username string, md *models.Task) *mongo.SingleResult {
	coll := DB.Collection("tasks")
	update := bson.M{
		"$set": md,
	}
	err := coll.FindOneAndUpdate(context.TODO(), bson.D{{"identifier", username}}, update)
	// _, err := coll.UpdateOne(context.TODO(), bson.D{{"control_number", cn}}, update)
	return err
}

func DeletePost(cn int) *mongo.SingleResult {
	coll := DB.Collection("posts")
	// _, err := coll.DeleteOne(context.TODO(), bson.D{{"control_number", cn}})
	err := coll.FindOneAndDelete(context.TODO(), bson.D{{"control_number", cn}})
	return err
}
