package controllers

import (
	pw "github.com/MarcosIgnacioo/playwright"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LogInUser(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	assigments, err := pw.ScrapMoodleAndClassroom(username, password)

	if err != nil {
		c.JSON(http.StatusTeapot, err)
		return
	}
	c.JSON(http.StatusOK, assigments)
}

func GetMoodleAssigments(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	assigments, err := pw.ScrapMoodle(username, password)

	if err != nil {
		c.JSON(http.StatusTeapot, err)
		return
	}
	c.JSON(http.StatusOK, assigments)
}

func GetClassroomAssigments(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	assigments, err := pw.ScrapClassroom(username, password)

	if err != nil {
		c.JSON(http.StatusTeapot, err)
		return
	}

	c.JSON(http.StatusOK, assigments)
}

func GetUserCredentials() {
}

func GetKardex() {
}

func GetCurricularMap() {
}
