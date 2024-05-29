package db

import (
	"context"
	"fmt"

	pw "github.com/MarcosIgnacioo/playwright"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAssigments(platform string, identifier string) (bson.M, error) {
	var result bson.M
	coll := DB.Collection(platform)
	err := coll.FindOne(context.TODO(), bson.D{{"name", identifier}}).
		Decode(&result)
	return result, err
}

func InsertClassRoom(identifier string, classRoom []interface{}) error {
	coll := DB.Collection("classroom")
	classRoomInfo := pw.NewClassRoomInfo(classRoom, identifier)
	res := coll.FindOneAndReplace(context.TODO(), bson.D{{"name", identifier}}, classRoomInfo)
	fmt.Println("de res", res)
	if res.Err() != nil {
		_, err := coll.InsertOne(context.TODO(), classRoomInfo)
		return err
	}
	return res.Err()
}

func InsertMoodle(identifier string, moodle []interface{}) error {
	coll := DB.Collection("moodle")
	moodleInfo := pw.NewMoodleInfo(moodle, identifier)
	res := coll.FindOneAndReplace(context.TODO(), bson.D{{"name", identifier}}, moodleInfo)
	if res.Err() != nil {
		_, err := coll.InsertOne(context.TODO(), moodleInfo)
		return err
	}
	return res.Err()
}

func InsertStudent(student *pw.StudentInfo) (*mongo.InsertOneResult, error) {
	_, checkStudent := GetStudent(student.ControlNumber)
	if checkStudent != nil {
		coll := DB.Collection("students")
		r, e := coll.InsertOne(context.TODO(), student)
		return r, e
	}
	return nil, nil
}

func UpdateStudent(cn int, st *pw.StudentInfo) error {
	coll := DB.Collection("students")
	update := bson.M{
		"$set": st,
	}
	_, err := coll.UpdateOne(context.TODO(), bson.D{{"control_number", cn}}, update)
	return err
}

func DeleteFromCollection(collection string, fieldName string, identifier string) error {
	coll := DB.Collection(collection)
	_, err := coll.DeleteOne(context.TODO(), bson.D{{fieldName, identifier}})
	return err
}

func DeleteStudent(identifier string) error {
	email := fmt.Sprintf(`%v@alu.uabcs.mx`, identifier)
	coll := DB.Collection("students")
	_, err := coll.DeleteOne(context.TODO(), bson.D{{"institutional_email", email}})
	return err
}

func GetStudent(cn int) (bson.M, error) {
	var result bson.M
	coll := DB.Collection("students")
	err := coll.FindOne(context.TODO(), bson.D{{"control_number", cn}}).
		Decode(&result)
	return result, err
}
