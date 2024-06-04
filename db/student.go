package db

// Ramtendo
//
// Francisco Alejandro Alcantar Aviles
// Marcos Ignacio Camacho Gonzalez
// Abraham Zumaya Manriquez
//
// package db
//
// Funciones CREATE, DELETE, UPDATE, GET de la base de datos

import (
	"context"
	"fmt"

	"github.com/MarcosIgnacioo/models"
	pw "github.com/MarcosIgnacioo/playwright"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Obtener las tareas del alumno, se le pasa el identificador del alumno junto a la plataforma de la que se va a obtener
func GetAssigments(platform string, identifier string) (bson.M, error) {
	var result bson.M
	coll := DB.Collection(platform)
	err := coll.FindOne(context.TODO(), bson.D{{"name", identifier}}).
		Decode(&result)
	return result, err
}

// Insertar las tareas del alumno en la colección de classroom
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

// Insertar las tareas del alumno en la colección de moodle
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

// Insertar los datos del estudiante a su colección correspondiente
func InsertStudent(student *pw.StudentInfo) (*mongo.InsertOneResult, error) {
	_, checkStudent := GetStudent(student.ControlNumber)
	if checkStudent != nil {
		coll := DB.Collection("students")
		r, e := coll.InsertOne(context.TODO(), student)
		return r, e
	}
	return nil, nil
}

// Actuializar los datos de un estudiante dado un número de control
func UpdateStudent(cn int, st *pw.StudentInfo) error {
	coll := DB.Collection("students")
	update := bson.M{
		"$set": st,
	}
	_, err := coll.UpdateOne(context.TODO(), bson.D{{"control_number", cn}}, update)
	return err
}

// Borrar los datos de un estudiante dado el identificador
func DeleteStudent(identifier string) error {
	email := fmt.Sprintf(`%v@alu.uabcs.mx`, identifier)
	coll := DB.Collection("students")
	_, err := coll.DeleteOne(context.TODO(), bson.D{{"institutional_email", email}})
	return err
}

// Obtener un estudiante dado un número de control
func GetStudent(cn int) (bson.M, error) {
	var result bson.M
	coll := DB.Collection("students")
	err := coll.FindOne(context.TODO(), bson.D{{"control_number", cn}}).
		Decode(&result)
	return result, err
}

// Función para insertar las `tasks` a la colección correspondiente con el identificador del usuario
func InsertTasks(tasks *models.Tasks) error {
	coll := DB.Collection("tasks")
	res := coll.FindOneAndReplace(context.TODO(), bson.D{{"identifier", tasks.Identifier}}, tasks)
	if res.Err() != nil {
		coll.InsertOne(context.TODO(), tasks)
	}
	return res.Err()
}

// Función para obtener las `tasks` del usuario
func GetTasks(identifier string) (bson.M, error) {
	var result bson.M
	coll := DB.Collection("tasks")
	coll.FindOne(context.TODO(), bson.D{{"identifier", identifier}}).
		Decode(&result)
	return result, nil
}

// Función para actualizar las `tasks` del usuario
func UpdateTask(username string, md *models.Task) *mongo.SingleResult {
	coll := DB.Collection("tasks")
	update := bson.M{
		"$set": md,
	}
	err := coll.FindOneAndUpdate(context.TODO(), bson.D{{"identifier", username}}, update)
	return err
}

// Función que sirve para borrar datos de una colección.
// Se pasa el nombre de la colección (string)
// El campo por el que se va a filtrar (string)
// Y el valor del filtro (string)
func DeleteFromCollection(collection string, fieldName string, identifier string) error {
	coll := DB.Collection(collection)
	_, err := coll.DeleteOne(context.TODO(), bson.D{{fieldName, identifier}})
	return err
}
