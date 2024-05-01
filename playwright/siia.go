package pw

import (
	"log"

	"github.com/playwright-community/playwright-go"
)

// La version asincrona
func SiiaScrapAsync(browser *playwright.Browser, username string, password string) *playwright.Page {
	siia, err := (*browser).NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	siia.Goto("https://siia.uabcs.mx/")
	siia.Locator("#ctl00_placeHolder_txtLogin").Fill(username)
	siia.Locator("#ctl00_placeHolder_txtPassword").Fill(password)
	siia.Locator("#ctl00_placeHolder_btnIniciarSesion").Click()
	return &siia
}

// La version sincrona

// kardex
// const subjects = document.querySelectorAll('tr')
// for (int i = 0; i<subjects.length; i++) {
// 	if (subjects[i].querySelectorAll('td').length == 9) {
// 		console.log(subjects[i].querySelector('td').innerText)
// 	}
// }

// mapa curricular
// Pueden haber null porque las materias que aun no han sido cursadas
// const semesters = document.querySelectorAll(".sce-semester")
// for(let i = 0; i<semesters.length; i++) {
//   const subjects = semesters[i].querySelectorAll(".sce-materia")
//   for (subject of subjects) {
//     console.log(subject.querySelector(".sce-name").innerText)
//     console.log(subject.querySelector(".sce-teacher").innerText)
//     console.log(subject.querySelector(".sce-period").innerText)
//     console.log(subject.querySelector(".sce-calification").innerText)
//   }
// }
// Hacer para las optativas
// sce-optative-program
