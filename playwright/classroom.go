package pw

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

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

func ClassroomScrap(context *playwright.BrowserContext, username string, password string, params ...string) (Result, error) {
	user := "1"
	if len(params) > 0 {
		fmt.Println("entraaa")
		user = params[0]
	} else {
		fmt.Println("no entra")
	}
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

	classroom.WaitForURL("https://classroom.google.com/")
	classroom.Goto("https://classroom.google.com/u/0/a/not-turned-in/all")
	expect.Locator(classroom.Locator("#LATER")).ToBeVisible()
	subjects, err := classroom.Locator("ol").All()
	scrappedAssigments := arraylist.NewArrayList(10)
	if err != nil {
		return nil, err
	}
	for _, subject := range subjects {
		assigments, err := subject.Locator("li").All()
		if err != nil {
			continue
		}
		for _, assigment := range assigments {
			link, _ := assigment.Locator("a").First().GetAttribute("href")
			bytedLink := []byte(link)
			bytedLink[3] = []byte(user)[0]
			link = string(bytedLink)
			link = fmt.Sprintf(`https://classroom.google.com%s`, link)
			ps, _ := assigment.Locator("p").All()
			title, _ := ps[0].InnerText()
			subjectName, _ := ps[1].InnerText()
			dueDate, _ := ps[2].InnerText()
			dateFormat := utils.DateFormat{}
			r, _ := regexp.Compile(`Fecha`)
			isValidDate := r.FindString(dueDate)

			if isValidDate == "" {
				var day, month, hour string
				year := fmt.Sprint(time.Now().Year())
				r, _ = regexp.Compile(`\d+:\d+`)
				hour = string(r.Find([]byte(dueDate)))
				if hour != "" {
					day = fmt.Sprint(time.Now().Day())
					time.Now().Local().Weekday()
					month = fmt.Sprint(time.Now().Month())
					month = utils.Months[month]
				} else {
					r, _ = regexp.Compile(`\d+ \w*`)
					dateData := r.FindAll([]byte(dueDate), -1)
					formattedDate := utils.Bytes_matrix_to_string(dateData)
					dateSplitted := strings.Split(formattedDate, " ")
					day = dateSplitted[0]
					month = dateSplitted[1]
					month = utils.Months[month]
					hour = "N/A"
				}
				fmt.Println(day, month, hour)
				dateFormat = utils.DateFormat{Day: day, Month: month, Year: year, Hour: hour}
			}
			scrappedAssigment := NewAssigment(subjectName, title, link, dateFormat)
			scrappedAssigments.Push(scrappedAssigment)
		}
	}
	// noDueDateAssigments := classroom.Locator("#NO_DUE_DATE").Last().Locator("ol").All()
	// thisWeekAssigments := classroom.Locator("#THIS_WEEK").Last().Locator("ol").All()
	// nextWeekAssigments := classroom.Locator("#NEXT_WEEK").Last().Locator("ol").All()
	// laterAssigments := classroom.Locator("#LATER").Last().Locator("ol").All()

	// XD
	// fmt.Println(noDueDateAssigments)
	// fmt.Println(thisWeekAssigments)
	// fmt.Println(nextWeekAssigments)

	fmt.Println(scrappedAssigments.GetArray()...)
	classroomAssigmentsArray := scrappedAssigments.GetArray()
	return NewClassRoomInfo(classroomAssigmentsArray), nil
}
