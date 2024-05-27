package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/MarcosIgnacioo/controllers"
	"github.com/MarcosIgnacioo/db"
	pw "github.com/MarcosIgnacioo/playwright"
	"github.com/gin-gonic/gin"
)

func main() {
	scraptest()
	db.Init()
	fmt.Println("BUILD GAMER")
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

	r.GET("/get-tasks", controllers.GetTasks)
	r.GET("/pancho", controllers.GetPapancho)
	r.POST("/save-tasks", controllers.PostTasks)

	r.GET("/ca", ramses)
	r.GET("/cb", ramses)
	// Esto va hasta el final!!!!! XD
	r.Run()
}

func scraptest() {
	context, _, _, err := pw.GenerateContext(true)

	if err != nil {
		fmt.Println(err)
		panic("ERROORROROORORORO")
	}

	page, err := (*context).NewPage()
	page2, err := (*context).NewPage()

	page.Goto("https://quotes.toscrape.com/")
	page2.Goto("https://es.wikipedia.org/wiki/Web_scraping")
	fmt.Println("wep")
	quote, _ := page.Locator(".quote span").First().InnerText()
	fmt.Println("ya hay quote")
	pw.Screenshot(&page, "CULO.png")
	title, _ := page2.Locator(".mw-page-title-main").First().InnerText()
	fmt.Println("ya hay title")
	fmt.Println(quote)
	fmt.Println(title)
	page2.Close()
	if len((*context).Pages()) == 1 {
		(*context).Close()
	} else {
		page.Close()
	}

	// sdfd, _ := pw.ScrapMoodleAndClassroom("marcosignc_21", "sopitasprecio")
	// browser, _, _ := pw.GenerateBrowser(false)
	// page, _ := (*browser).NewPage()
	// page.Goto("https://es.wikipedia.org/wiki/Web_scraping")
	// time.Sleep(2 * time.Second)
	// page.Context().NewPage()
	// res := pw.Cronos(false, browser, "marcosignc_21", "sopitasprecio", pw.MoodleScrap)
	// pw.Cronos(false, browser, "marcosignc_21", "sopitasprecio", pw.MoodleScrap)
	// pw.Cronos(false, browser, "marcosignc_21", "sopitasprecio", pw.MoodleScrap)
	// pw.Cronos(false, browser, "marcosignc_21", "sopitasprecio", pw.MoodleScrap)
	// fmt.Println("this should go on")
	// fmt.Println("and then crash")
	// fmt.Println(res)
	// fmt.Println(sdfd.Moodle...)
	// fmt.Println(sdfd.ClassRoom...)
	// username := "marcosignc_21"
	// password := "sopitasprecio"
	// browser, _, _ := pw.GenerateBrowser(false)
	// siia, _ := pw.SiiaLogin(browser, username, password)
	// newPage2, _ := (*siia).Context().NewPage()
	// pw.KardexScrap(siia)
	// session, _ := (*browser).NewContext()
	// page, _ := session.NewPage()
	// page.Goto("netflix.com")

	// pw.CurricularMapScrap(&newPage2)
	// pw.Testing("marcosignc_21", "sopitasprecio")
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
