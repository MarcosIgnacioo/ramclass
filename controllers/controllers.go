package controllers

import (
	"net/http"

	"github.com/MarcosIgnacioo/db"
	pw "github.com/MarcosIgnacioo/playwright"
	"github.com/gin-gonic/gin"
)

func LogInUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	scrappedInfo, err := pw.FullScrap(username, password)

	// XDDDD
	go db.InsertStudent(scrappedInfo.(*pw.ScrappedInfo).Student.(*pw.StudentInfo))
	go db.InsertClassRoom(scrappedInfo.(*pw.ScrappedInfo).Student.(*pw.StudentInfo), scrappedInfo.(*pw.ScrappedInfo).ClassRoom)
	go db.InsertMoodle(scrappedInfo.(*pw.ScrappedInfo).Student.(*pw.StudentInfo), scrappedInfo.(*pw.ScrappedInfo).Moodle)
	go db.InsertKardex(username, scrappedInfo.(*pw.ScrappedInfo).GPA.(*pw.GPA).GPA, scrappedInfo.(*pw.ScrappedInfo).Kardex)
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
	go db.InsertMoodle(assigments.(*pw.ScrappedInfo).Student.(*pw.StudentInfo), assigments.(*pw.ScrappedInfo).Moodle)

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
	go db.InsertClassRoom(assigments.(*pw.ScrappedInfo).Student.(*pw.StudentInfo), assigments.(*pw.ScrappedInfo).ClassRoom)

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

	go db.InsertStudent(credentials.(*pw.ScrappedInfo).Student.(*pw.StudentInfo))

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

	go db.InsertKardex(username, subjects.(*pw.ScrappedInfo).GPA.(*pw.GPA).GPA, subjects.(*pw.ScrappedInfo).Kardex)

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
	go db.InsertCurricularMap(username, subjects.(*pw.ScrappedInfo).CurricularMap)
	if err != nil {
		c.JSON(http.StatusTeapot, err)
		return
	}
	c.JSON(http.StatusOK, subjects)
}
