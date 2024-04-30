package pw

import (
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/MarcosIgnacioo/arraylist"
	"github.com/MarcosIgnacioo/utils"
	"github.com/playwright-community/playwright-go"
)

var expect = playwright.NewPlaywrightAssertions(10000)
var await = playwright.NewPlaywrightAssertions(500)

func ClassroomScrapNoChannel(browser *playwright.Browser, username string, password string) (*[]interface{}, error) {
	classroom, err := (*browser).NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	classroom.Goto("https://accounts.google.com/ServiceLogin?continue=https%3A%2F%2Fclassroom.google.com&passive=true")
	expect.Locator(classroom.Locator("#identifierId")).ToBeVisible()
	classroom.Locator("#identifierId").Fill(fmt.Sprintf("%v@alu.uabcs.mx", username))
	// Por alguna razon en headless el form no ha sido actualizado por lo que si estyoy haciendolo con el navegador activado pues tengo que usar el nth pero si lo hago haciendo el modo headless con el GetByText
	// classroom.Locator("button").Nth(2).Click()
	classroom.GetByText("Next").Click()
	classroom.Locator("#username").Fill(username)
	classroom.Locator("#password").Fill(password)
	classroom.Locator("input").Nth(2).Click()
	// TODO agregar un checador de que si hay error con la contrase;a retorne error jejejej
	expect.Locator(classroom.Locator(".hrUpcomingAssignmentGroup > a").Last()).ToBeVisible()
	classes, _ := classroom.Locator("li:has(.hrUpcomingAssignmentGroup)").All()
	scrappedAssigments := arraylist.NewArrayList(10)
	for _, class := range classes {
		assigment := class.Locator(".hrUpcomingAssignmentGroup > a").First()
		subject, _ := class.Locator("h2 a div").First().TextContent()
		title, _ := assigment.GetAttribute("aria-label")
		link, _ := assigment.GetAttribute("href")
		link = fmt.Sprintf("https://classroom.google.com%v", link)
		scrappedAssigment := NewAssigment(subject, title, link, utils.DateFormat{})
		scrappedAssigments.Push(scrappedAssigment)
	}
	classroomAssigmentsArray := scrappedAssigments.GetArray()
	return &classroomAssigmentsArray, nil
}

func ClassroomScrap(browser *playwright.Browser, username string, password string, cs chan []interface{}) {
	classroom, err := (*browser).NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	classroom.Goto("https://accounts.google.com/ServiceLogin?continue=https%3A%2F%2Fclassroom.google.com&passive=true")
	expect.Locator(classroom.Locator("#identifierId")).ToBeVisible()
	classroom.Locator("#identifierId").Fill(fmt.Sprintf("%v@alu.uabcs.mx", username))
	// Por alguna razon en headless el form no ha sido actualizado por lo que si estyoy haciendolo con el navegador activado pues tengo que usar el nth pero si lo hago haciendo el modo headless con el GetByText
	// classroom.Locator("button").Nth(2).Click()
	classroom.GetByText("Next").Click()
	classroom.Locator("#username").Fill(username)
	classroom.Locator("#password").Fill(password)
	classroom.Locator("input").Nth(2).Click()
	expect.Locator(classroom.Locator(".hrUpcomingAssignmentGroup > a").Last()).ToBeVisible()
	classes, _ := classroom.Locator("li:has(.hrUpcomingAssignmentGroup)").All()
	scrappedAssigments := arraylist.NewArrayList(10)
	for _, class := range classes {
		assigment := class.Locator(".hrUpcomingAssignmentGroup > a").First()
		subject, _ := class.Locator("h2 a div").First().TextContent()
		title, _ := assigment.GetAttribute("aria-label")
		link, _ := assigment.GetAttribute("href")
		link = fmt.Sprintf("https://classroom.google.com%v", link)
		scrappedAssigment := NewAssigment(subject, title, link, utils.DateFormat{})
		scrappedAssigments.Push(scrappedAssigment)
	}
	cs <- scrappedAssigments.GetArray()
}

func MoodleScrap(browser *playwright.Browser, username string, password string) ([]interface{}, error) {
	moodle, err := (*browser).NewPage()

	if err != nil {
		log.Fatalf("could not create moodle: %v", err)
	}

	if _, err = moodle.Goto("https://enlinea2024-1.uabcs.mx/login/"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	moodle.Locator("#username").Fill(username)
	moodle.Locator("#password").Fill(password)
	moodle.Locator("#loginbtn").Click()
	url := moodle.URL()
	if url != "https://enlinea2024-1.uabcs.mx/my/" {
		err := errors.New("Credenciales incorrectas")
		return nil, err
	}

	expect.Locator(moodle.Locator(".multiline")).ToBeVisible()
	tabContent, _ := moodle.Locator(".event-name-container").All()

	subjects := arraylist.NewArrayList(10)
	for _, v := range tabContent {
		classSubject, _ := v.Locator("small").InnerText()
		anchorTag := v.Locator("a").First()
		dueDate, _ := anchorTag.GetAttribute("aria-label")
		r, _ := regexp.Compile(`\d \w* \w* \w* \d*, \d*:\d*`)
		date := r.FindAll([]byte(dueDate), -1)
		dateFormated := utils.CreateDate(date)

		assigmentTitle, assError := anchorTag.InnerText()
		if assError != nil {
			assigmentTitle = "No hay titulo"
		}
		link, linkErr := anchorTag.GetAttribute("href")
		if linkErr != nil {
			link = "No hay link"
		}
		subjects.Push(NewAssigment(classSubject, assigmentTitle, link, *dateFormated))
	}
	return subjects.GetArray(), nil
}

// Pasar esto a su propio package

func ScrapMoodleAndClassroom(username string, password string) (*ScrappedInfo, *LoginError) {
	// TODO: Crear un package con variables globales (Expect)

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	// playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(false)}
	//                                vv
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(true)})

	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	cs := make(chan []interface{})

	go ClassroomScrap(&browser, username, password, cs)
	ms, logErr := MoodleScrap(&browser, username, password)

	if logErr != nil {
		if err = browser.Close(); err != nil {
			log.Fatalf("could not close browser: %v", err)
		}
		if err = pw.Stop(); err != nil {
			log.Fatalf("could not stop Playwright: %v", err)
		}
		return nil, NewLoginError(logErr.Error())
	}

	moodleArray := arraylist.NewArrayList(10)
	classroomArray := <-cs
	for _, v := range ms {
		moodleArray.Push(v)
	}

	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}

	mArr := moodleArray.GetArray()
	return NewScrappedInfo(mArr, classroomArray), nil
}

func ScrapMoodle(username string, password string) (*[]interface{}, *LoginError) {

	pw, err := playwright.Run()

	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(true)})

	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	ms, logErr := MoodleScrap(&browser, username, password)

	if logErr != nil {
		if err = browser.Close(); err != nil {
			log.Fatalf("could not close browser: %v", err)
		}
		if err = pw.Stop(); err != nil {
			log.Fatalf("could not stop Playwright: %v", err)
		}
		return nil, NewLoginError(logErr.Error())
	}

	moodleArray := arraylist.NewArrayList(10)
	for _, v := range ms {
		moodleArray.Push(v)
	}

	mArr := moodleArray.GetArray()
	return &mArr, nil
}

func ScrapClassroom(username string, password string) (*[]interface{}, *LoginError) {

	pw, err := playwright.Run()

	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(true)})

	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	ca, logErr := ClassroomScrapNoChannel(&browser, username, password)

	if logErr != nil {
		if err = browser.Close(); err != nil {
			log.Fatalf("could not close browser: %v", err)
		}
		if err = pw.Stop(); err != nil {
			log.Fatalf("could not stop Playwright: %v", err)
		}
		return nil, NewLoginError(logErr.Error())
	}

	return ca, nil
}

func screenshot(page playwright.Page, name string) {
	page.Screenshot((playwright.PageScreenshotOptions{
		Path: playwright.String(name),
	}))
}

func logError(e *error) {
	if e != nil {
		log.Fatalf("error ocurred: %v", e)
	}
}
func closingError(logErr *error, browser *playwright.Browser, pw *playwright.Playwright) *LoginError {
	if logErr != nil {
		if err := (*browser).Close(); err != nil {
			log.Fatalf("could not close browser: %v", err)
		}
		if err := pw.Stop(); err != nil {
			log.Fatalf("could not stop Playwright: %v", err)
		}
		return NewLoginError((*logErr).Error())
	}
	return nil
}
