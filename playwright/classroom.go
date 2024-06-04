package pw

// Ramtendo
//
// Francisco Alejandro Alcantar Aviles
// Marcos Ignacio Camacho Gonzalez
// Abraham Zumaya Manriquez
//
// package pw
// Aquí es donde se encuentra la funcionalidad del web scrapping del classroom.

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

// # Función para realizr el scrappeo de classroom
// # context  *playwright.BrowserContext
// # username string
// # password string
// # params   ...string
// Se le pasa un puntero del contexto del navegador (para poder mantener todo dentro de una misma instancia de Chromium), el nombre de uusario y su contraseña, junto a un parámetro extra, el cual puede no ir, dicho parámetro se utiliza para saber el número de usuario al que corresponden los links de las tareas, por defecto es 1
func ClassroomScrap(context *playwright.BrowserContext, username string, password string, params ...string) (Result, error) {
	user := "1"
	if len(params) > 0 {
		user = params[0]
	}
	classroom, err := (*context).NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	classroom.Goto("https://accounts.google.com/ServiceLogin?continue=https%3A%2F%2Fclassroom.google.com&passive=true")
	expect.Locator(classroom.Locator("#identifierId")).ToBeVisible()
	email := fmt.Sprintf("%v@alu.uabcs.mx", username)
	classroom.Locator("#identifierId").Fill(email)

	// Por alguna razon en headless el form no ha sido actualizado por lo que si estyoy haciendolo con el navegador activado pues tengo que usar el nth pero si lo hago haciendo el modo headless con el GetByText
	// classroom.Locator("button").Nth(2).Click()
	// error handling que no se si deba d kitar jiiji

	expect.Locator(classroom.Locator("#identifierId")).ToHaveValue(email)
	classroom.GetByText("Next").Click()
	classroom.Locator("#username").Fill(username)
	expect.Locator(classroom.Locator("#username")).ToHaveValue(username)
	classroom.Locator("#password").Fill(password)
	expect.Locator(classroom.Locator("#password")).ToHaveValue(password)
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
				r, _ = regexp.Compile(`\p{L}+, \d+:\d+`)
				hour = string(r.Find([]byte(dueDate)))
				if hour != "" {
					dateSplitted := strings.Split(hour, ",")
					day = dateSplitted[0]
					month = ""
					hour = dateSplitted[1]
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
				dateFormat = utils.DateFormat{Day: day, Month: month, Year: year, Hour: hour}
			}
			scrappedAssigment := NewAssigment(subjectName, title, link, dateFormat)
			scrappedAssigments.Push(scrappedAssigment)
		}
	}
	classroomAssigmentsArray := scrappedAssigments.GetArray()
	return NewClassRoomInfo(classroomAssigmentsArray), nil
}
