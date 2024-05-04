package controllers

import (
	"net/http"

	pw "github.com/MarcosIgnacioo/playwright"
	"github.com/gin-gonic/gin"
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

func GetUserCredentials(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	subjects, err := pw.SiiaInit("credentials", username, password)

	if err != nil {
		c.JSON(http.StatusTeapot, err)
		return
	}

	c.JSON(http.StatusOK, subjects)
}

func GetKardex(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	subjects, err := pw.SiiaInit("kardex", username, password)

	if err != nil {
		c.JSON(http.StatusTeapot, err)
		return
	}

	c.JSON(http.StatusOK, subjects)

}

func GetCurricularMap(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	subjects, err := pw.SiiaInit("map", username, password)
	if err != nil {
		c.JSON(http.StatusTeapot, err)
		return
	}
	c.JSON(http.StatusOK, subjects)
}
