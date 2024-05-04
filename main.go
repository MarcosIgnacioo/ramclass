package main

import (
	"fmt"
	"net/http"

	"github.com/MarcosIgnacioo/controllers"
	pw "github.com/MarcosIgnacioo/playwright"
	"github.com/gin-gonic/gin"
)

func main() {
	// scraptest()
	server()
	// Compute(false)
	// Compute(false)
	// Compute(false)
	// Compute(false)
	// Compute(false)
	// Compute(false)
	// Compute(false)
	// Compute(true)
	// Compute(false)
}
func exec() (bool, error) {
	response := false
	return response, nil
}

func Compute(sync bool) {
	if !sync {
		type result struct {
			value bool
			err   error
		}
		c := make(chan result)
		go func() {
			var r result
			r.value, r.err = exec()
			c <- r
		}()
		fmt.Println(<-c)
	} else {
		res, err := exec()
		fmt.Println(res)
		fmt.Println(err)
	}
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
	r.GET("/kardex-form", ramses)
	r.GET("/credentials-form", ramses)
	r.GET("/curricular-form", ramses)
	r.GET("/test", ramses)

	r.POST("/login-user", controllers.LogInUser)

	r.POST("/classroom", controllers.GetClassroomAssigments)
	r.POST("/moodle", controllers.GetMoodleAssigments)

	r.POST("/kardex", controllers.GetKardex)
	r.POST("/curricular", controllers.GetCurricularMap)
	r.POST("/credentials", controllers.GetUserCredentials)

	// Esto va hasta el final!!!!! XD
	r.Run()
}
func ramses(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
