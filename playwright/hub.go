package pw

// Estas serian las unicas funciones que van a ser publicas en el package
const (
	headless = true
)

func FullScrap(username string, password string) (Result, *Error) {
	context, browser, playwright, err := GenerateContext(headless)
	if err != nil {
		return nil, (*Error)(NewError(err))
	}

	credentialsResults, loginError := CredentialsScrap(context, username, password)

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

	// Getting the results from scrapping into channels

	// fmt.Println(<-classRoomChannel)

	moodleResults := <-moodleChannel

	classRoomResults := <-classRoomChannel

	kardexResults := <-kardexChannel

	curricularResults := <-curricularChannel

	// aqui se termina  el scrapping
	closingErr := CloseScrapper(playwright, browser, context)
	if closingErr != nil {
		return nil, (*Error)(NewError(closingErr))
	}

	GPA := NewGPA(kardexResults.(*Kardex).GPA)

	return NewScrappedInfo(moodleResults.GetResult(), classRoomResults.GetResult(), kardexResults.GetResult(), curricularResults.GetResult(), credentialsResults, GPA), (*Error)(NewError(nil))
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
