package db

import (
	"context"
	pw "github.com/MarcosIgnacioo/playwright"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertStudent(student *pw.StudentInfo) (*mongo.InsertOneResult, error) {
	coll := DB.Collection("students")
	r, e := coll.InsertOne(context.TODO(), student)
	return r, e
}

func UpdateStudent(cn int, st *pw.StudentInfo) error {
	coll := DB.Collection("students")
	update := bson.M{
		"$set": st,
	}
	_, err := coll.UpdateOne(context.TODO(), bson.D{{"control_number", cn}}, update)
	return err
}

func DeleteStudent(cn int) error {
	coll := DB.Collection("students")
	_, err := coll.DeleteOne(context.TODO(), bson.D{{"control_number", cn}})
	return err
}

func GetStudent(cn int) (bson.M, error) {
	var result bson.M
	coll := DB.Collection("students")
	coll.FindOne(context.TODO(), bson.D{{"control_number", cn}}).
		Decode(&result)
	return result, nil
}
