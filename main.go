package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/MarcosIgnacioo/controllers"
	pw "github.com/MarcosIgnacioo/playwright"
	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("wep")
	// db.Init()
	// db.Mongo()
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
func assertErrorToNilf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func assertEqual(expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		panic(fmt.Sprintf("%v does not equal %v", actual, expected))
	}
}

const todoName = "Bake a cake"

func scraptest() {
	context, _, _, err := pw.GenerateContext(false)

	assertErrorToNilf("could not create context: %w", err)
	page, err := (*context).NewPage()
	page2, err := (*context).NewPage()

	assertErrorToNilf("could not create page: %w", err)
	page.Goto("https://quotes.toscrape.com/")
	page2.Goto("https://es.wikipedia.org/wiki/Web_scraping")
	fmt.Println("wep")
	quote, _ := page.Locator(".quote span").First().InnerText()
	fmt.Println("ya hay quote")
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

func server() {
	r := gin.Default()
	r.Static("/public/assets", "./public/assets/")
	r.LoadHTMLGlob("views/*")
	r.GET("/", ramses)
	r.GET("/sign-in", ramses)
	r.GET("/student", ramses)
	r.GET("/home", ramses)
	r.GET("/my-kardex", ramses)
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
	r.POST("/curricular", controllers.GetCurricularMap)
	r.POST("/credentials", controllers.GetUserCredentials)

	r.GET("/ca", ramses)
	r.GET("/cb", ramses)
	// Esto va hasta el final!!!!! XD
	r.Run()
}
func ramses(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
