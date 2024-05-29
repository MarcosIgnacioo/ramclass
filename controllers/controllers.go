package controllers

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
		scrappedInfo.(*pw.ScrappedInfo).Tasks = bson.M{}
	}

	scrappedInfo.(*pw.ScrappedInfo).Calendar = tasks

	c.JSON(http.StatusOK, scrappedInfo)
}

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

func GetClassroomAssigments(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	assigments, err := pw.Classroom(username, password)
	go db.InsertClassRoom(username, assigments.(*pw.ClassRoomInfo).ClassRoom)

	if err != nil {
		c.JSON(http.StatusTeapot, err)
		return
	}

	c.JSON(http.StatusOK, assigments)
}

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

func SaveTasks(c *gin.Context) {
	username := c.PostForm("identifier")
	tasks := c.PostForm("tasks")
	fmt.Println(username, tasks)
	// if err != nil {
	// 	c.JSON(http.StatusTeapot, err)
	// 	return
	// }
	//
	c.JSON(http.StatusOK, tasks)
}

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
		c.JSON(http.StatusOK, bson.M{})
	}
}

func DeleteStudent(c *gin.Context) {
	id, err := c.GetQuery("identifier")
	deleteError := db.DeleteStudent(id)
	if deleteError != nil {
		c.JSON(http.StatusConflict, deleteError)
		return
	}
	collections := []string{"classroom", "moodle", "kardex", "curricular_map", "tasks"}
	for _, collection := range collections {
		deleteError = db.DeleteFromCollection(collection, "name", id)
		if deleteError != nil {
			c.JSON(http.StatusConflict, deleteError)
			return
		}
	}
	if err != true {
		c.JSON(http.StatusConflict, err)
		return
	}
	c.JSON(http.StatusOK, bson.M{"first": id})
}
