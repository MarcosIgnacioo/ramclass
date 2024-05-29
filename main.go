package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/MarcosIgnacioo/controllers"
	"github.com/MarcosIgnacioo/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	server()
}

func server() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	os.Setenv("HOME", "/home/marcig/")
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.LoadHTMLGlob(dir + "/views/*")
	r.Static("/public/assets", dir+"/public/assets/")

	r.GET("/", ramses)
	r.GET("/sign-in", ramses)
	r.GET("/student", ramses)
	r.GET("/home", ramses)
	r.GET("/my-kardex", ramses)
	r.GET("/my-curricular-map", ramses)
	r.GET("/todo", ramses)
	r.GET("/calendar", ramses)
	r.GET("/settings", ramses)
	r.GET("/faq", ramses)

	r.GET("/terms-of-privacy", termsOfPrivacy)
	r.GET("/terms-of-service", termsOfService)

	r.GET("/create-account", ramses)
	r.GET("/classroom-form", ramses)
	r.GET("/moodle-form", ramses)
	r.GET("/kardex-form", ramses)
	r.GET("/credentials-form", ramses)
	r.GET("/curricular-form", ramses)
	r.GET("/test", ramses)

	r.POST("/login-user", controllers.LogInUser)
	r.POST("/create-account", controllers.LogInUser)

	r.POST("/classroom", controllers.GetClassroomAssigments)
	r.POST("/moodle", controllers.GetMoodleAssigments)

	r.POST("/kardex", controllers.GetKardex)
	r.POST("/curricular-map", controllers.GetCurricularMap)
	r.POST("/credentials", controllers.GetUserCredentials)

	r.DELETE("/student", controllers.DeleteStudent)

	r.GET("/get-tasks", controllers.GetTasks)
	r.POST("/save-tasks", controllers.PostTasks)

	r.GET("/ca", ramses)
	r.GET("/cb", ramses)
	// Esto va hasta el final!!!!! XD
	r.Run()
}

func ramses(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func termsOfPrivacy(c *gin.Context) {
	terms, _ := db.GetTermsOfPrivacy()
	c.JSON(http.StatusOK, terms["terms_of_privacy"])
}

func termsOfService(c *gin.Context) {
	terms, _ := db.GetTermsOfService()
	c.JSON(http.StatusOK, terms["terms_of_service"])
}
