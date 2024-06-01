package pw

import (
	"fmt"
	"log"

	"github.com/MarcosIgnacioo/arraylist"
	"github.com/MarcosIgnacioo/utils"
	"github.com/playwright-community/playwright-go"
)

func ClassroomScrapAsync(browser *playwright.Browser, username string, password string, cs chan []interface{}) {
	classroom, err := (*browser).NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	classroom.Goto("https://accounts.google.com/ServiceLogin?continue=https%3A%2F%2Fclassroom.google.com&passive=true")
	expect.Locator(classroom.Locator("#identifierId")).ToBeVisible()
	classroom.Locator("#identifierId").Fill(fmt.Sprintf("%v@alu.uabcs.mx", username))
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

func ClassroomScrap(context *playwright.BrowserContext, username string, password string) (Result, error) {
	classroom, err := (*context).NewPage()
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

	screenshot(&classroom, "wtf2.png")
	classroom.Close()
	classroomAssigmentsArray := scrappedAssigments.GetArray()
	return NewClassRoomInfo(classroomAssigmentsArray), nil
}
