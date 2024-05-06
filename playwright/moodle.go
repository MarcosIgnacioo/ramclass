package pw

import (
	"errors"
	"log"
	"regexp"

	"github.com/MarcosIgnacioo/arraylist"
	"github.com/MarcosIgnacioo/utils"
	"github.com/playwright-community/playwright-go"
)

func MoodleScrap(context *playwright.BrowserContext, username string, password string) (Result, error) {
	moodle, err := (*context).NewPage()
	//
	if err != nil {
		log.Fatalf("could not create moodle: %v", err)
	}
	//
	if _, err = moodle.Goto("https://enlinea2024-1.uabcs.mx/login/"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	expect.Locator(moodle.Locator("#loginbtn")).ToBeVisible()
	moodle.Locator("#username").Fill(username)
	moodle.Locator("#password").Fill(password)
	moodle.Locator("#loginbtn").Click()
	url := moodle.URL()
	if url != "https://enlinea2024-1.uabcs.mx/my/" {
		err := errors.New("Credenciales incorrectas")
		return nil, err
	}
	//
	expect.Locator(moodle.Locator(".multiline")).ToBeVisible()
	tabContent, _ := moodle.Locator(".event-name-container").All()
	//
	subjects := arraylist.NewArrayList(10)
	for _, v := range tabContent {
		classSubject, _ := v.Locator("small").InnerText()
		anchorTag := v.Locator("a").First()
		dueDate, _ := anchorTag.GetAttribute("aria-label")
		r, _ := regexp.Compile(`\d \w* \w* \w* \d*, \d*:\d*`)
		date := r.FindAll([]byte(dueDate), -1)
		dateFormated := utils.CreateDate(date)
		//
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
	moodle.Close()
	return NewMoodleInfo(subjects.GetArray()), nil
}
