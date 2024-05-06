package pw

import "fmt"

// Estas serian las unicas funciones que van a ser publicas en el package
const (
	headless = false
)

func FullScrap(username string, password string) (Result, *Error) {
	context, browser, playwright, err := GenerateContext(headless)
	if err != nil {
		return nil, (*Error)(NewError(err))
	}

	credentialsAsync := CreateAsyncScrapping(CredentialsScrap, context, username, password)
	credentialsChannel := make(chan Result)
	credentialsErrChannel := make(chan error)
	go credentialsAsync(credentialsChannel, credentialsErrChannel)

	classRoomAsync := CreateAsyncScrapping(ClassroomScrap, context, username, password)
	classRoomChannel := make(chan Result)
	classRoomErrChannel := make(chan error)
	go classRoomAsync(classRoomChannel, classRoomErrChannel)

	kardexAsync := CreateAsyncScrapping(KardexScrap, context, username, password)
	kardexChannel := make(chan Result)
	kardexErrChannel := make(chan error)
	go kardexAsync(kardexChannel, kardexErrChannel)

	testc := make(chan Result)
	testerr := make(chan error)
	go kardexAsync(testc, testerr)

	curricularMapAsync := CreateAsyncScrapping(CurricularMapScrap, context, username, password)
	curricularChannel := make(chan Result)
	curricularErrChannel := make(chan error)
	go curricularMapAsync(curricularChannel, curricularErrChannel)

	moodleAsync := CreateAsyncScrapping(MoodleScrap, context, username, password)
	moodleChannel := make(chan Result)
	moodleErrChannel := make(chan error)
	go moodleAsync(moodleChannel, moodleErrChannel)

	// Getting the results from scrapping into channels
	classRoomResults := <-classRoomChannel
	classRoomErr := <-classRoomErrChannel

	kardexResults := <-kardexChannel
	kardexErr := <-kardexErrChannel

	curricularResults := <-curricularChannel
	curricularErr := <-curricularErrChannel

	moodleResults := <-moodleChannel
	moodleErr := <-moodleErrChannel

	credentialsResults := <-credentialsChannel
	credentialsErr := <-credentialsErrChannel

	fmt.Println(kardexErr)
	fmt.Println(classRoomErr)
	fmt.Println(curricularErr)
	fmt.Println(moodleErr)
	// fmt.Println(credentialsResults.GetResult()...)
	fmt.Println(credentialsErr)

	// aqui se termina  el scrapping
	closingErr := CloseScrapper(playwright, browser, context)
	if closingErr != nil {
		fmt.Println("wep")
		return nil, (*Error)(NewError(closingErr))
	}
	// fmt.Println(classRoomErr)
	return NewScrappedInfo(moodleResults.GetResult(), classRoomResults.GetResult(), kardexResults.GetResult(), curricularResults.GetResult(), credentialsResults.GetResult()), (*Error)(NewError(classRoomErr))
	// return NewScrappedInfo(nil, nil, nil, nil, credentialsResults.GetResult()), nil
}

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

func Classroom(username string, password string) (Result, *Error) {
	context, browser, playwright, err := GenerateContext(headless)
	if err != nil {
		return nil, (*Error)(NewError(err))
	}
	res, err := ClassroomScrap(context, username, password)
	closingErr := CloseScrapper(playwright, browser, context)
	if closingErr != nil {
		return nil, (*Error)(NewError(err))
	}
	(*context).Close()
	return res, nil
}

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
