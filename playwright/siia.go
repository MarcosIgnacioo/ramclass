package pw

// Ramtendo
//
// Francisco Alejandro Alcantar Aviles
// Marcos Ignacio Camacho Gonzalez
// Abraham Zumaya Manriquez
//
// package pw
// Aquí es donde se encuentra la funcionalidad del web scrapping.
// Específicamente el del siia

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

// # Función para iniciar sesión en la página del siia
// # siia     *playwright.Page
// # url      string
// # username string
// # password string
// El parametro de url sirve para saber si ya se ha iniciado sesión, es decir, primero intenta ir a una de las vistas del siia, por ejemplo el kardex; En el caso de que ya haya iniciado sesión antes en alguno de las otras go routines del scrapping del siia, la página estará en donde tiene que estar, si no es asi, se le redirigirá a una vista de login, en la que se iniciará la sesión.
func siiaLogin(siia *playwright.Page, url string, username string, password string) error {
	(*siia).Goto(url)
	// Si redirige bien a la url que queremos significa que ya habia iniciado sesion por lo que simplemente hacemos un early return
	if (*siia).URL() == url {
		return nil
	}
	expect.Locator((*siia).Locator("#ctl00_ContentPlaceHolder1_initLinkButton1")).ToBeVisible()
	(*siia).Locator("#ctl00_ContentPlaceHolder1_loginTextBox").Fill(username)
	(*siia).Locator("#ctl00_ContentPlaceHolder1_passwordTextBox").Fill(password)
	(*siia).Locator("#ctl00_ContentPlaceHolder1_initLinkButton1").Click()

	if (*siia).URL() != url {
		return errors.New("Credenciales incorrectas")
	}
	return nil
}

// # Función para realizar el scrapping del kardex
// # context  *playwright.BrowserContext
// # username string
// # password string
// # params   ...string
// Realiza el scrapping necesario para obtener los datos del kardex
func KardexScrap(context *playwright.BrowserContext, username string, password string, params ...string) (Result, error) {
	kardexUrl := "https://siia.uabcs.mx/siia2019/alumnos/kardex.aspx?gr=alumno"
	siia, _ := (*context).NewPage()
	siia.Goto(kardexUrl)
	err := siiaLogin(&siia, kardexUrl, username, password)
	if err != nil {
		return nil, err
	}
	subjectsArrayList := arraylist.NewArrayList(100)
	//
	expect.Locator(siia.Locator("tbody tr")).ToBeVisible()
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
	gpaString, _ := siia.Locator("tfoot tr .number").InnerText()
	gpa, _ := strconv.Atoi(gpaString)
	siia.Close()
	return NewKardex(gpa, subjectsArrayList.GetArray()), nil
}

// # Función para hacer el scrapping del mapa curricular
// # context  *playwright.BrowserContext
// # username string
// # password string
// # params   ...string
// Realiza el scrapping necesario para obtener los datos del mapa curricular
func CurricularMapScrap(context *playwright.BrowserContext, username string, password string, params ...string) (Result, error) {
	curricularUrl := "https://siia.uabcs.mx/siia2019/alumnos/mapaacademico.aspx?gr=alumno"
	subjectsArrayList := arraylist.NewArrayList(100)

	siia, _ := (*context).NewPage()
	err := siiaLogin(&siia, curricularUrl, username, password)
	expect.Locator(siia.Locator(".sce-semester")).ToBeVisible()
	semesters, err := siia.Locator(".sce-semester").All()
	if err != nil {
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

// # Función para hacer el scrapping de las credenciales del estudiante
// # context  *playwright.BrowserContext
// # username string
// # password string
// # params   ...string
// Realiza el scrapping necesario para obtener los datos de las credenciales del usuario
func CredentialsScrap(context *playwright.BrowserContext, username string, password string) (Result, error) {
	siia, _ := (*context).NewPage()
	credentialsUrl := "https://siia.uabcs.mx/siia2019/alumnos/credenciales.aspx?gr=alumno"
	err := siiaLogin(&siia, credentialsUrl, username, password)
	if err != nil {
		return nil, err
	}
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
