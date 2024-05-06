package pw

// CERRAR PLAYWRIGHT CADA VEZ QUE SE HAGA UN SCRAPPPPP !!!
import (
	"github.com/MarcosIgnacioo/arraylist"
	"github.com/playwright-community/playwright-go"
	"log"
)

func ScrapMoodleAndClassroom(username string, password string) (*ScrappedInfo, *LoginError) {

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(false)})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	context, err := browser.NewContext()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	cs := make(chan []interface{})

	ctx, err := browser.NewContext()

	if err != nil {
		return nil, NewError(err)
	}

	go SiiaLogin(&ctx, username, password)
	go ClassroomScrapAsync(&browser, username, password, cs)
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
	classroomArray := <-cs
	for _, v := range ms.GetResult() {
		moodleArray.Push(v)
	}

	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}

	mArr := moodleArray.GetArray()
	return NewScrappedInfo(mArr, classroomArray, nil, nil, nil), nil
}

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
