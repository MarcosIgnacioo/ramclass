package main

import (
	"fmt"
	"net/http"

	"github.com/MarcosIgnacioo/controllers"
	pw "github.com/MarcosIgnacioo/playwright"
	"github.com/MarcosIgnacioo/types"
	"github.com/gin-gonic/gin"
)

func main() {
	server()
}

func scraptest() {
	user := types.User{Username: "", Password: ""}
	pw.ScrapMoodleAndClassroom("", "")
	fmt.Println(user)
}

func server() {
	r := gin.Default()
	r.Static("/public/assets", "./public/assets/")
	r.LoadHTMLGlob("views/*")
	r.GET("/", ramses)
	r.GET("/classroom-form", ramses)
	r.GET("/moodle-form", ramses)
	r.GET("/test", ramses)

	r.POST("/login-user", controllers.LogInUser)
	r.POST("/classroom", controllers.GetClassroomAssigments)
	r.POST("/moodle", controllers.GetMoodleAssigments)

	// Esto va hasta el final!!!!! XD
	r.Run()
}
func ramses(c *gin.Context) {
	// c.JSON(http.StatusOK, "asdf")
	c.HTML(http.StatusOK, "index.html", nil)
}
