package pw

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/MarcosIgnacioo/classroomapi"
	"github.com/MarcosIgnacioo/types"
	"github.com/playwright-community/playwright-go"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/classroom/v1"
)

func ClassroomAPIAuth(browser *playwright.Browser, authURL string, username string, password string) (*string, error) {
	classroom, err := (*browser).NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	// TODO manejo de errores djdjfajdjf
	classroom.Goto(authURL)
	classroom.Locator("#identifierId").Fill(fmt.Sprintf("%v@alu.uabcs.mx", username))
	classroom.GetByText("Next").Click()
	classroom.Locator("#username").Fill(username)
	classroom.Locator("#password").Fill(password)
	classroom.Locator("input").Nth(2).Click()
	expect.Locator(classroom.Locator("button").Last()).ToBeVisible()
	classroom.Locator("button").Last().Click()
	expect.Locator(classroom.GetByText("github")).ToBeVisible()
	classroom.Locator("button").Last().Click()
	classroom.WaitForURL("https://marcosignacioo.github.io/*")
	url, err := url.Parse(classroom.URL())
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	APIcode := url.Query().Get("code")
	classroom.Close()
	return &APIcode, nil
}

func StartScrapping(user types.User) {
	// TODO: Crear un package con variables globales (Expect)
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	// playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(false)}
	//                                vv
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(false)})

	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	// cs := make(chan []interface{})

	b, err := os.ReadFile("credentials.json")
	config, err := google.ConfigFromJSON(b, classroom.ClassroomCoursesReadonlyScope)
	authURL := classroomapi.GetTokenURL(config)
	tokenAPI, err := ClassroomAPIAuth(&browser, *authURL, user.Username, user.Password)
	fmt.Println(tokenAPI)
	if err != nil {
		return
	}
	classroomapi.GetOAuthURL(tokenAPI)
}
