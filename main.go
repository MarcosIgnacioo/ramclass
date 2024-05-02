package main

import (
	"net/http"

	"github.com/MarcosIgnacioo/controllers"
	pw "github.com/MarcosIgnacioo/playwright"
	"github.com/gin-gonic/gin"
)

func main() {
	scraptest()
}

func scraptest() {
	// sdfd, _ := pw.ScrapMoodleAndClassroom("marcosignc_21", "sopitasprecio")
	// fmt.Println(sdfd.Moodle...)
	// fmt.Println(sdfd.ClassRoom...)
	pw.SiiaInit("map", "marcosignc_21", "sopitasprecio")
	// pw.Testing("marcosignc_21", "sopitasprecio")
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
