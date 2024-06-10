package controllers

// Ramtendo
//
// Francisco Alejandro Alcantar Aviles
// Marcos Ignacio Camacho Gonzalez
// Abraham Zumaya Manriquez
//
// package controllers
// Se definen las funciones controladoras de los endpoints de nuestra api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MarcosIgnacioo/db"
	"github.com/MarcosIgnacioo/models"
	pw "github.com/MarcosIgnacioo/playwright"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// # Login Handler
//
// # c *gin.Context
//
// Controlador para manejar el inicio de sesión del usuario, se obtiene el nombre de usuario del formulario junto a su contraseña y se le pasa a la funcion que hace el web scrapping. Después se insertan a la base de datos los resultados obtenidos a sus respectivas colecciones.

func LogInUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	scrappedInfo, err := pw.FullScrap(username, password)

	if err.ErrorMessage != nil {
		c.JSON(http.StatusTeapot, err)
		return
	}

	// GO routines go brrrrr
	go db.InsertStudent(scrappedInfo.(*pw.ScrappedInfo).Student.(*pw.StudentInfo))
	go db.InsertClassRoom(username, scrappedInfo.(*pw.ScrappedInfo).ClassRoom)
	go db.InsertMoodle(username, scrappedInfo.(*pw.ScrappedInfo).Moodle)
	go db.InsertKardex(username, scrappedInfo.(*pw.ScrappedInfo).GPA, scrappedInfo.(*pw.ScrappedInfo).Kardex)
	go db.InsertCurricularMap(username, scrappedInfo.(*pw.ScrappedInfo).CurricularMap)

	tasks, errTasks := db.GetTasks(username)

	if errTasks != nil {
		fmt.Println(errTasks)
		return
	}

	if tasks["tasks"] != nil {
		scrappedInfo.(*pw.ScrappedInfo).Tasks = tasks["tasks"].(primitive.M)
	} else {
		scrappedInfo.(*pw.ScrappedInfo).Tasks = bson.M{
			"Lunes":     []bson.M{{"is_done": false, "task_description": ""}},
			"Martes":    []bson.M{{"is_done": false, "task_description": ""}},
			"Miércoles": []bson.M{{"is_done": false, "task_description": ""}},
			"Jueves":    []bson.M{{"is_done": false, "task_description": ""}},
			"Viernes":   []bson.M{{"is_done": false, "task_description": ""}},
			"Sábado":    []bson.M{{"is_done": false, "task_description": ""}},
			"Domingo":   []bson.M{{"is_done": false, "task_description": ""}},
		}
	}

	scrappedInfo.(*pw.ScrappedInfo).Calendar = tasks

	c.JSON(http.StatusOK, scrappedInfo)
}

// # Moodle Assigments Handler
//
// # c *gin.Context
//
// Sirve para obtener unicamente las tareas de moodle, se obtiene el nombre de usuario del formulario junto a su contraseña y se le pasa a la función que hace el web scrapping, y luego se inserta en la colección de la base de datos correspondiente.

func GetMoodleAssigments(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	assigments, err := pw.Moodle(username, password)
	go db.InsertMoodle(username, assigments.(*pw.MoodleInfo).Moodle)

	if err != nil {
		c.JSON(http.StatusTeapot, err)
		return
	}
	c.JSON(http.StatusOK, assigments)
}

// # Classroom Assigments Handler
//
// # c *gin.Context
//
// Sirve para obtener unicamente las tareas de classroom, se obtiene el nombre de usuario del formulario junto a su contraseña y se le pasa a la función que hace el web scrapping, y luego se inserta en la colección de la base de datos correspondiente.

func GetClassroomAssigments(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	classroomId := c.PostForm("classroom-id")
	assigments, err := pw.Classroom(username, password, classroomId)
	go db.InsertClassRoom(username, assigments.(*pw.ClassRoomInfo).ClassRoom)

	if err != nil {
		c.JSON(http.StatusTeapot, err)
		return
	}

	c.JSON(http.StatusOK, assigments)
}

// # User Credentials Handler
//
// # c *gin.Context
//
// Sirve para obtener unicamente las credenciales del estudiante, se obtiene el nombre de usuario del formulario junto a su contraseña y se le pasa a la función que hace el web scrapping, y luego se inserta en la colección de la base de datos correspondiente.

func GetUserCredentials(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	credentials, err := pw.StudentCredential(username, password)

	go db.InsertStudent(credentials.(*pw.StudentInfo))

	if err != nil {
		c.JSON(http.StatusTeapot, err)
		return
	}

	c.JSON(http.StatusOK, credentials)
}

// # Kardex Handler
//
// # c *gin.Context
//
// Sirve para obtener unicamente el kardex del estudiante, se obtiene el nombre de usuario del formulario junto a su contraseña y se le pasa a la función que hace el web scrapping, y luego se inserta en la colección de la base de datos correspondiente.

func GetKardex(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	subjects, err := pw.Grades(username, password)

	go db.InsertKardex(username, subjects.(*pw.Kardex).GPA, subjects.(*pw.Kardex).Kardex)

	if err != nil {
		c.JSON(http.StatusTeapot, err)
		return
	}

	c.JSON(http.StatusOK, subjects)

}

// # Curriculara Map Handler
//
// # c *gin.Context
//
// Sirve para obtener unicamente las materias del mapa curricular del alumno, se obtiene el nombre de usuario del formulario junto a su contraseña y se le pasa a la función que hace el web scrapping, y luego se inserta en la colección de la base de datos correspondiente.

func GetCurricularMap(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	subjects, err := pw.CareerSubjects(username, password)
	go db.InsertCurricularMap(username, subjects.(*pw.CurricularMap).CurricularMap)
	if err != nil {
		c.JSON(http.StatusTeapot, err)
		return
	}
	c.JSON(http.StatusOK, subjects)
}

// # Post Tasks Handler
//
// # c *gin.Context
//
// Sirve para guardar las `tasks` que mande el usuario por medio del formulario en la colección correspondiente en la base de datos, identificada por el nombre del usuario

func PostTasks(c *gin.Context) {
	tasksString := c.PostForm("tasks")
	var tasks models.Tasks
	json.Unmarshal([]byte(tasksString), &tasks)
	err := db.InsertTasks(&tasks)
	if err != nil {
		c.JSON(http.StatusConflict, err)
		return
	}
	c.JSON(http.StatusOK, "ok")
}

// # Get Tasks Handler
//
// # c *gin.Context
//
// Sirve para obtener las `tasks` del usuario que tenga guardadas en la base de datos, se obtienen a partir del identificador del estudiante, el cual manda por medio de los query params. En caso de no tener `tasks` enviamos un objeto de `tasks` vacío.

func GetTasks(c *gin.Context) {
	identifier, err := c.GetQuery("identifier")
	fmt.Println(identifier)
	if err != true {
		c.JSON(http.StatusConflict, err)
		return
	}
	res, errDb := db.GetTasks(identifier)
	if errDb != nil {
		c.JSON(http.StatusConflict, err)
		return
	}

	if res["tasks"] != nil {
		c.JSON(http.StatusOK, res["tasks"])
	} else {
		c.JSON(http.StatusOK, bson.M{
			"Lunes":     []bson.M{{"is_done": false, "task_description": ""}},
			"Martes":    []bson.M{{"is_done": false, "task_description": ""}},
			"Miércoles": []bson.M{{"is_done": false, "task_description": ""}},
			"Jueves":    []bson.M{{"is_done": false, "task_description": ""}},
			"Viernes":   []bson.M{{"is_done": false, "task_description": ""}},
			"Sábado":    []bson.M{{"is_done": false, "task_description": ""}},
			"Domingo":   []bson.M{{"is_done": false, "task_description": ""}},
		})
	}
}

// # Delete Student Handler
//
// # c *gin.Context
//
// Sirve para borrar toda la información del estudiante que ha sido guardada en nuestra base de datos

func DeleteStudent(c *gin.Context) {
	id, err := c.GetQuery("identifier")

	if err != true {
		c.JSON(http.StatusConflict, "There is no identifier")
		return
	}

	collections := []string{"classroom", "moodle", "kardex", "curricular_map"}
	for _, collection := range collections {
		db.DeleteFromCollection(collection, "name", id)
	}
	db.DeleteFromCollection("tasks", "identifier", id)
	db.DeleteFromCollection("students", "institutional_email", fmt.Sprintf("%s@alu.uabcs.mx", id))

	c.JSON(http.StatusOK, bson.M{"first": id})
}
