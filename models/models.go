package models

// Ramtendo
//
// Francisco Alejandro Alcantar Aviles
// Marcos Ignacio Camacho Gonzalez
// Abraham Zumaya Manriquez
//
// package models
//
// Modelos necesarios para la inserción de datos a las colecciones en mongodb

// Struct para guardar las tareas
type Tasks struct {
	Identifier string      `bson:"identifier" json:"identifier"`
	Tasks      interface{} `bson:"tasks" json:"tasks"`
}

// Struct para una sola tarea
type Task struct {
	Day  string `bson:"day" json:"day"`
	Task string `bson:"task" json:"task"`
}
