package db

import (
	"context"

	pw "github.com/MarcosIgnacioo/playwright"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Generic struct {
	Student    *pw.StudentInfo
	Assigments *[]interface{} `json:"assigments"`
}

func NewGeneric(s *pw.StudentInfo, a *[]interface{}) *Generic {
	return &Generic{Student: s, Assigments: a}
}

func InsertClassRoom(student *pw.StudentInfo, classRoom []interface{}) (*mongo.InsertOneResult, error) {
	coll := DB.Collection("classroom")
	r, e := coll.InsertOne(context.TODO(), NewGeneric(student, &classRoom))
	return r, e
}
func InsertMoodle(student *pw.StudentInfo, moodle []interface{}) (*mongo.InsertOneResult, error) {
	coll := DB.Collection("moodle")
	r, e := coll.InsertOne(context.TODO(), NewGeneric(student, &moodle))
	return r, e
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

func DeleteStudent(cn int) error {
	coll := DB.Collection("students")
	_, err := coll.DeleteOne(context.TODO(), bson.D{{"control_number", cn}})
	return err
}

func GetStudent(cn int) (bson.M, error) {
	var result bson.M
	coll := DB.Collection("students")
	err := coll.FindOne(context.TODO(), bson.D{{"control_number", cn}}).
		Decode(&result)
	return result, err
}
