package pw

// Ramtendo
//
// Francisco Alejandro Alcantar Aviles
// Marcos Ignacio Camacho Gonzalez
// Abraham Zumaya Manriquez
//
// package pw
// Aquí es donde se encuentra la funcionalidad del web scrapping.

import (
	"log"

	"github.com/MarcosIgnacioo/arraylist"
	"github.com/playwright-community/playwright-go"
)

// Función para generar una instancia de Chromium en la que se puede compartir las pestañas entre scrapping.
//
// headless bool
//
// false -> abre el navegador con gui
//
// true -> no abre el navegador lo hace por websockets(oalgoasi)
func GenerateContext(headless bool) (*playwright.BrowserContext, *playwright.Browser, *playwright.Playwright, error) {
	//
	pw, err := playwright.Run()
	//
	if err != nil {
		return nil, nil, nil, err
	}
	//
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(headless)})
	//
	if err != nil {
		return nil, nil, pw, err
	}
	context, err := browser.NewContext()

	if err != nil {
		return nil, nil, pw, err
	}
	return &context, &browser, pw, err
	//
}

// Función para inicializar el proceso del scrappeo de moodle
// username string
// password string
func ScrapMoodle(username string, password string) (*[]interface{}, *LoginError) {

	pw, err := playwright.Run()

	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(true)})

	if err != nil {
		return nil, NewError(err)
	}

	context, err := browser.NewContext()
	if err != nil {
		return nil, NewError(err)
	}

	ms, logErr := MoodleScrap(&context, username, password)

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
	for _, v := range ms.GetResult() {
		moodleArray.Push(v)
	}

	mArr := moodleArray.GetArray()
	return &mArr, nil
}

// Función para inicializar el proceso del scrappeo de classrooom
// username string
// password string
func ScrapClassroom(username string, password string) (Result, *LoginError) {

	pw, err := playwright.Run()

	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(true)})

	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	context, err := browser.NewContext()

	ca, logErr := ClassroomScrap(&context, username, password)

	context.Close()
	CloseScrapper(pw, &browser, &context)

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
