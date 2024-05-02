package pw

import (
	"fmt"
	"log"
	"strconv"

	"github.com/MarcosIgnacioo/arraylist"
	"github.com/playwright-community/playwright-go"
)

// TOdo agregar en parametros el nombre de usuario y contra
func SiiaInit(action string, username string, password string) {
	browser, _, err := GenerateBrowser(false)
	if err != nil {
		log.Fatalf("error opening the browser")
	}
	siia := SiiaLogin(browser, username, password)
	switch action {
	case "kardex":
		KardexScrap(siia)
	case "map":
		CurricularMapScrap(siia)
	}
}

// La version sincrona
func SiiaLogin(browser *playwright.Browser, username string, password string) *playwright.Page {
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

func KardexScrap(siia *playwright.Page) (Result, error) {
	subjectsArrayList := arraylist.NewArrayList(100)
	(*siia).Goto("https://siia.uabcs.mx/siia2019/alumnos/kardex.aspx?gr=alumno&op=kardex")
	//
	rows, err := (*siia).Locator("tbody tr").All()
	if err != nil {
		log.Fatalf("rows not loaded %v", err)
		return nil, err
	}
	//
	var semester int
	for _, row := range rows {
		columns, err := row.Locator("td").All()
		if err != nil {
			log.Fatal("columns not loaded")
			return nil, err
		}
		if len(columns) == 9 {
			semester++
			columns = columns[1:]
		}
		subjectName, _ := columns[0].InnerText()
		group, _ := columns[1].InnerText()
		turn, _ := columns[2].InnerText()
		period, _ := columns[3].InnerText()
		gradeString, _ := columns[4].InnerText()
		grade, _ := strconv.Atoi(gradeString)
		state, _ := columns[5].InnerText()
		subjectType, _ := columns[6].InnerText()
		teacher, _ := columns[7].InnerText()
		subjectsArrayList.Enqueue(NewSubject(semester, subjectName, group, turn, period, grade, state, subjectType, teacher))
	}
	return NewKardex(subjectsArrayList.GetArray()), nil
}

func CurricularMapScrap(siia *playwright.Page) (Result, error) {
	subjectsArrayList := arraylist.NewArrayList(100)
	(*siia).Goto("https://siia.uabcs.mx/siia2019/alumnos/mapaacademico.aspx?gr=alumno&op=mapa_alum")
	//
	semesters, err := (*siia).Locator(".sce-semester").All()
	semesters = semesters[0 : len(semesters)-1]
	if err != nil {
		log.Fatalf("rows not loaded %v", err)
		return nil, err
	}
	fmt.Println(len(semesters))
	// de un for no se esta saliendo quien sabe por q
	for i, semester := range semesters {
		semesterNo := i + 1
		fmt.Println(semesterNo)
		subjects, _ := semester.Locator(".sce-materia").All()
		for _, subject := range subjects {
			// por alguna razon se corta en la parte de ingenieria de software 2
			subjectName, _ := subject.Locator(".materia-name span").InnerText()
			period, _ := subject.Locator(".sce-period").InnerText()
			// Manejar el caso de que no haya calificacion (Es decir que aun no se cursa esa materia)
			// ya 8)
			// por alguna razon no esta agarrando el resto de materias que deberian de aparecer con el 255 y los valores default
			grade := 255
			state := ""
			subjectType := ""
			teacher := ""
			credits := 255
			if period != "" {
				gradeString, _ := subject.Locator(".sce-calification").InnerText()
				grade, _ = strconv.Atoi(gradeString)
				creditsString, _ := subject.Locator(".sce-creditos").InnerText()
				credits, _ = strconv.Atoi(creditsString)
				state, _ = subject.Locator(".sce-status").InnerText()
				subjectType, _ = subject.Locator(".sce-tipo-ex").InnerText()
				teacher, _ = subject.Locator(".sce-teacher").InnerText()
			}
			subjectsArrayList.Enqueue(NewCurricularSubject(semesterNo, subjectName, period, grade, state, credits, subjectType, teacher))
		}
	}
	cm := NewCurricularMap(subjectsArrayList.GetArray())
	fmt.Println("wep")
	fmt.Println(cm.GetResult()...)
	fmt.Println("wep")
	return cm, nil
}

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
