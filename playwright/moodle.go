package pw

// Ramtendo
//
// Francisco Alejandro Alcantar Aviles
// Marcos Ignacio Camacho Gonzalez
// Abraham Zumaya Manriquez
//
// package pw
// Aquí es donde se encuentra la funcionalidad del web scrapping de moodle.

import (
	"log"
	"os"
	"regexp"

	"github.com/MarcosIgnacioo/arraylist"
	"github.com/MarcosIgnacioo/utils"
	"github.com/playwright-community/playwright-go"
)

// # Función para realizr el scrappeo de moodle
// # context  *playwright.BrowserContext
// # username string
// # password string
// # params   ...string (En este caso no se usa, pero se tiene que poner en la firma de la función para poder pasarsela a la que convierte cualquier función en asíncrona debido a que todas deben de tener la misma firma)
func MoodleScrap(context *playwright.BrowserContext, username string, password string, params ...string) (Result, error) {
	moodle, err := (*context).NewPage()
	//
	if err != nil {
		log.Fatalf("could not create moodle: %v", err)
	}
	//

	if _, err = moodle.Goto(os.Getenv("MOODLE_URL")); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	expect.Locator(moodle.Locator("#loginbtn")).ToBeVisible()
	moodle.Locator("#username").Fill(username)
	// agregue esto para evitar que se le trabe la cola a esto
	expect.Locator(moodle.Locator("#loginbtn")).ToHaveValue(username)

	moodle.Locator("#password").Fill(password)
	// agregue esto para evitar que se le trabe la cola a esto
	expect.Locator(moodle.Locator("#password")).ToHaveValue(password)

	moodle.Locator("#loginbtn").Click()
	moodle.WaitForURL(os.Getenv("MOODLE_HOME_URL"))
	// if url != "https://enlinea2024-1.uabcs.mx/my/" {
	// 	err := errors.New("Credenciales incorrectas")
	// 	return nil, err
	// }
	//
	expect.Locator(moodle.Locator(".multiline")).ToBeVisible()
	tabContent, _ := moodle.Locator(".event-name-container").All()
	//
	subjects := arraylist.NewArrayList(10)
	for _, v := range tabContent {
		classSubject, _ := v.Locator("small").InnerText()
		anchorTag := v.Locator("a").First()
		dueDate, _ := anchorTag.GetAttribute("aria-label")
		r, _ := regexp.Compile(`\d+ \w* \w* \w* \d*, \d*:\d*`)
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
