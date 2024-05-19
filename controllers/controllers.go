package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MarcosIgnacioo/db"
	"github.com/MarcosIgnacioo/models"
	pw "github.com/MarcosIgnacioo/playwright"
	"github.com/gin-gonic/gin"
)

func LogInUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	scrappedInfo, err := pw.FullScrap(username, password)

	// GO routines go brrrrr
	go db.InsertStudent(scrappedInfo.(*pw.ScrappedInfo).Student.(*pw.StudentInfo))
	go db.InsertClassRoom(username, scrappedInfo.(*pw.ScrappedInfo).ClassRoom)
	go db.InsertMoodle(username, scrappedInfo.(*pw.ScrappedInfo).Moodle)
	go db.InsertKardex(username, scrappedInfo.(*pw.ScrappedInfo).GPA, scrappedInfo.(*pw.ScrappedInfo).Kardex)
	go db.InsertCurricularMap(username, scrappedInfo.(*pw.ScrappedInfo).CurricularMap)

	if err.ErrorMessage != nil {
		c.JSON(http.StatusTeapot, err)
		return
	}

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
	json.Unmarshal([]byte(tasksString), tasks)
	err := db.InsertTasks(tasks.Identifier, &tasks)

	if err != nil {
		c.JSON(http.StatusTeapot, err)
		return
	}

	c.JSON(http.StatusOK, "ok")
}
