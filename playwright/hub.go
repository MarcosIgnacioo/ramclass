package pw

// Ramtendo
//
// Francisco Alejandro Alcantar Aviles
// Marcos Ignacio Camacho Gonzalez
// Abraham Zumaya Manriquez
//
// package pw
// Aquí es donde se encuentra la funcionalidad del web scrapping.

import "fmt"

// Estas serian las unicas funciones que van a ser publicas en el package
const (
	headless = true
)

// Realiza el scrappeo entero de todas los sitios que se pueden scrappear
func FullScrap(username string, password string, params ...string) (Result, *Error) {
	context, browser, playwright, err := GenerateContext(headless)
	if err != nil {
		return nil, (*Error)(NewError(err))
	}

	credentialsResults, loginError := CredentialsScrap(context, username, password)
	fmt.Println(credentialsResults)
	fmt.Println(loginError)

	if loginError != nil {
		CloseScrapper(playwright, browser, context)
		return nil, (*Error)(NewError(loginError))
	}

	classRoomAsync := CreateAsyncScrapping(ClassroomScrap, context, username, password)
	classRoomChannel := make(chan Result)
	classRoomErrChannel := make(chan error)
	go classRoomAsync(classRoomChannel, classRoomErrChannel)

	moodleAsync := CreateAsyncScrapping(MoodleScrap, context, username, password)
	moodleChannel := make(chan Result)
	moodleErrChannel := make(chan error)
	go moodleAsync(moodleChannel, moodleErrChannel)

	kardexAsync := CreateAsyncScrapping(KardexScrap, context, username, password)
	kardexChannel := make(chan Result)
	kardexErrChannel := make(chan error)
	go kardexAsync(kardexChannel, kardexErrChannel)

	curricularMapAsync := CreateAsyncScrapping(CurricularMapScrap, context, username, password)
	curricularChannel := make(chan Result)
	curricularErrChannel := make(chan error)
	go curricularMapAsync(curricularChannel, curricularErrChannel)

	moodleResults := <-moodleChannel

	classRoomResults := <-classRoomChannel

	kardexResults := <-kardexChannel

	curricularResults := <-curricularChannel

	closingErr := CloseScrapper(playwright, browser, context)
	if closingErr != nil {
		return nil, (*Error)(NewError(closingErr))
	}

	GPA := kardexResults.(*Kardex).GPA

	return NewScrappedInfo(moodleResults.GetResult(), classRoomResults.GetResult(), kardexResults.GetResult(), curricularResults.GetResult(), credentialsResults, GPA), (*Error)(NewError(nil))
}

// Se inicializa el scrappeo del moodle
func Moodle(username string, password string) (Result, *Error) {
	context, browser, playwright, err := GenerateContext(headless)
	if err != nil {
		return nil, (*Error)(NewError(err))
	}
	res, err := MoodleScrap(context, username, password)
	closingErr := CloseScrapper(playwright, browser, context)
	if closingErr != nil {
		return nil, (*Error)(NewError(err))
	}
	return res, nil
}

// Se inicializa el scrappeo de classroom
func Classroom(username string, password string, params ...string) (Result, *Error) {
	user := "1"
	if len(params) > 0 {
		user = params[0]
	}
	context, browser, playwright, err := GenerateContext(headless)
	if err != nil {
		return nil, (*Error)(NewError(err))
	}
	res, err := ClassroomScrap(context, username, password, user)
	closingErr := CloseScrapper(playwright, browser, context)
	if closingErr != nil {
		return nil, (*Error)(NewError(err))
	}
	(*context).Close()
	return res, nil
}

// Se inicializa el scrappeo del kardex
func Grades(username string, password string) (Result, *Error) {
	context, browser, playwright, err := GenerateContext(headless)
	if err != nil {
		return nil, (*Error)(NewError(err))
	}
	res, errSiia := KardexScrap(context, username, password)
	if errSiia != nil {
		return nil, (*Error)(NewError(errSiia))
	}
	closingErr := CloseScrapper(playwright, browser, context)
	if closingErr != nil {
		return nil, (*Error)(NewError(err))
	}
	(*context).Close()
	return res, nil
}

// Se inicializa el scrappeo del mapa curricular
func CareerSubjects(username string, password string) (Result, *Error) {
	context, browser, playwright, err := GenerateContext(headless)
	if err != nil {
		return nil, (*Error)(NewError(err))
	}
	res, errSiia := CurricularMapScrap(context, username, password)
	if errSiia != nil {
		return nil, (*Error)(NewError(errSiia))
	}
	closingErr := CloseScrapper(playwright, browser, context)
	if closingErr != nil {
		return nil, (*Error)(NewError(err))
	}
	return res, nil
}

// Se inicializa el scrappeo de las credenciales del estudiante
func StudentCredential(username string, password string) (Result, *Error) {
	context, browser, playwright, err := GenerateContext(headless)
	if err != nil {
		return nil, (*Error)(NewError(err))
	}

	res, errSiia := CredentialsScrap(context, username, password)
	if errSiia != nil {
		return nil, (*Error)(NewError(errSiia))
	}
	closingErr := CloseScrapper(playwright, browser, context)
	if closingErr != nil {
		return nil, (*Error)(NewError(err))
	}
	return res, nil
}
