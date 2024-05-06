package pw

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/MarcosIgnacioo/arraylist"
	"github.com/playwright-community/playwright-go"
)

const (
	APPROVED = "sce-student-course-approved"
	CURRENT  = "sce-student-course-current"
	NONE     = "sce-tipo-regular"
	OPTATIVE = "sce-tipo-optativa"
)

func SiiaLogin(context *playwright.BrowserContext, username string, password string) error {
	siia, err := (*context).NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	siia.Goto("https://siia.uabcs.mx/")
	expect.Locator(siia.Locator("#ctl00_placeHolder_btnIniciarSesion")).ToBeVisible()
	siia.Locator("#ctl00_placeHolder_txtLogin").Fill(username)
	siia.Locator("#ctl00_placeHolder_txtPassword").Fill(password)
	siia.Locator("#ctl00_placeHolder_btnIniciarSesion").Click()
	if siia.URL() == "https://siia.uabcs.mx/" {
		return errors.New("Credenciales incorrectas")
	}
	siia.Close()
	return nil
}

func KardexScrap(context *playwright.BrowserContext, username string, password string) (Result, error) {
	err := SiiaLogin(context, username, password)
	if err != nil {
		return nil, err
	}
	subjectsArrayList := arraylist.NewArrayList(100)
	siia, _ := (*context).NewPage()
	siia.Goto("https://siia.uabcs.mx/siia2019/alumnos/kardex.aspx?gr=alumno&op=kardex")
	//
	rows, err := siia.Locator("tbody tr").All()
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
	siia.Close()
	return NewKardex(subjectsArrayList.GetArray()), nil
}

func CurricularMapScrap(context *playwright.BrowserContext, username string, password string) (Result, error) {
	err := SiiaLogin(context, username, password)
	if err != nil {
		return nil, err
	}
	subjectsArrayList := arraylist.NewArrayList(100)

	siia, _ := (*context).NewPage()
	siia.Goto("https://siia.uabcs.mx/siia2019/alumnos/mapaacademico.aspx?gr=alumno&op=mapa_alum")
	//
	semesters, err := siia.Locator(".sce-semester").All()
	if err != nil {
		log.Fatalf("rows not loaded %v", err)
		return nil, err
	}
	//
	for i, semester := range semesters {
		semesterNo := i + 1
		subjects, _ := semester.Locator(".sce-materia").All()
		//
		for _, subject := range subjects {
			subjectName, _ := subject.Locator(".materia-name span").InnerText()
			classList, _ := subject.GetAttribute("class")
			classListArray := strings.Split(classList, " ")
			subjectState := classListArray[len(classListArray)-1]
			//
			var grade interface{}
			var credits interface{}
			var state interface{}
			var subjectType interface{}
			var teacher interface{}
			var period interface{}
			//
			creditsString, _ := subject.Locator(".sce-creditos").InnerText()
			credits, _ = strconv.Atoi(creditsString)
			//
			skip := false
			//
			if subjectState == APPROVED || subjectState == CURRENT {
				grade = nil
				period, _ = subject.Locator(".sce-period").InnerText()
				state, _ = subject.Locator(".sce-status").InnerText()
				subjectType, _ = subject.Locator(".sce-tipo-ex").InnerText()
				teacher, _ = subject.Locator(".sce-teacher").InnerText()
			}
			//
			if subjectState != APPROVED {
				skip = true
			}
			//
			if subjectState == OPTATIVE {
				semesterNo = 0
			}
			if skip {
				subjectsArrayList.Enqueue(NewCurricularSubject(semesterNo, subjectName, period, grade, state, credits, subjectType, teacher))
				continue
			}

			//
			gradeString, _ := subject.Locator(".sce-calification").InnerText()
			grade, _ = strconv.Atoi(gradeString)
			//
			subjectsArrayList.Enqueue(NewCurricularSubject(semesterNo, subjectName, period, grade, state, credits, subjectType, teacher))
		}
	}
	siia.Close()
	cm := NewCurricularMap(subjectsArrayList.GetArray())
	return cm, nil
}

// XD
func CredentialsScrap(context *playwright.BrowserContext, username string, password string) (Result, error) {
	err := SiiaLogin(context, username, password)
	if err != nil {
		return nil, err
	}
	siia, _ := (*context).NewPage()
	siia.Goto("https://siia.uabcs.mx/siia2019/alumnos/credenciales.aspx?gr=alumno&op=photocrede")
	expect.Locator(siia.Locator("#ctl00_contentPlaceHolder_alumnosFormView_AlumnoFieldset_AlumnoIDLabel")).ToBeVisible()
	controlNumber, _ := siia.Locator("#ctl00_contentPlaceHolder_alumnosFormView_AlumnoFieldset_AlumnoIDLabel").InnerText()
	studentId, _ := strconv.Atoi(controlNumber)
	studentName, _ := siia.Locator("#ctl00_contentPlaceHolder_alumnosFormView_AlumnoFieldset_Label1").InnerText()
	// El correo ya lo tenemos xd asi que podriamos omitirlo
	studentEmail, _ := siia.Locator("#ctl00_contentPlaceHolder_alumnosFormView_AlumnoFieldset_EmailLabel").InnerText()
	// el campu pues tambien
	studentCampus, _ := siia.Locator("#ctl00_contentPlaceHolder_alumnosFormView_AlumnoFieldset_NombreCampusLabel").InnerText()
	studentCareer, _ := siia.Locator("#ctl00_contentPlaceHolder_alumnosFormView_AlumnoFieldset_NombreCarreraLabel").InnerText()
	studentPeriod, _ := siia.Locator("#ctl00_contentPlaceHolder_alumnosFormView_AlumnoFieldset_PeriodoLabel").InnerText()
	semesterString, _ := siia.Locator("#ctl00_contentPlaceHolder_alumnosFormView_AlumnoFieldset_SemestreLabel").InnerText()
	currentSemester, _ := strconv.Atoi(semesterString)
	studentGroup, _ := siia.Locator("#ctl00_contentPlaceHolder_alumnosFormView_AlumnoFieldset_GrupoLabel").InnerText()
	studentTurn, _ := siia.Locator("#ctl00_contentPlaceHolder_alumnosFormView_AlumnoFieldset_TurnoLabel").InnerText()
	studentState, _ := siia.Locator("#ctl00_contentPlaceHolder_alumnosFormView_AlumnoFieldset_StatusAlumnoLabel").InnerText()

	siia.Close()
	return NewStudentInfo(studentId, studentName, studentEmail, studentCampus, studentCareer, studentPeriod, currentSemester, studentGroup, studentTurn, studentState), nil
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
