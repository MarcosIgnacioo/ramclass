package pw

import (
	"fmt"
	"log"

	"github.com/MarcosIgnacioo/arraylist"
	"github.com/playwright-community/playwright-go"
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

	cs := make(chan []interface{})

	go SiiaScrapAsync(&browser, username, password)
	go ClassroomScrapAsync(&browser, username, password, cs)
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
	return NewScrappedInfo(mArr, classroomArray), nil
}

func GenerateBrowser(headless bool) (*playwright.Browser, *playwright.Playwright, error) {
	//
	pw, err := playwright.Run()
	//
	if err != nil {
		return nil, nil, err
	}
	//
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(headless)})
	//
	if err != nil {
		return nil, pw, err
	}
	return &browser, pw, err
	//
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

	ca, logErr := ClassroomScrap(&browser, username, password)

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

func Testing(username string, password string) {
	browser, pw, err := GenerateBrowser(false)
	if err != nil {
		panic(err)
	}

	cr := make(chan Result)
	ce := make(chan error)
	asyncClassroom := CreateAsyncScrapping(ClassroomScrap, browser, username, password)
	go asyncClassroom(cr, ce)
	mood, err, _, _ := ChronoScrap(MoodleScrap, true, browser, username, password)
	err = CloseScrapper(pw, browser)
	if err != nil {
		panic(err)
	}
	classroomAssigments := <-cr
	fmt.Println(classroomAssigments.GetResult()...)
	fmt.Println(mood.GetResult()...)
}
