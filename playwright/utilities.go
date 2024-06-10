package pw

// Ramtendo
//
// Francisco Alejandro Alcantar Aviles
// Marcos Ignacio Camacho Gonzalez
// Abraham Zumaya Manriquez
//
// package pw
// Aqui se encuntran todas las funciones de utilidades que usamos para el scrapping

import (
	"errors"
	"fmt"

	"github.com/playwright-community/playwright-go"
)

// Variables que sirven para realizar assertions en el momento del scrapping
var expect = playwright.NewPlaywrightAssertions(10000)
var loginSiiaExpect = playwright.NewPlaywrightAssertions(1000)
var await = playwright.NewPlaywrightAssertions(500)

/*
Funcion que convierte en asincrona/go routine a la que se le pase con la firma indicada

scrapFn: func(*playwright.BrowserContext, string, string, ...string) (Result, error)

Definirá si la función sera ejecutada de manera síncrona

	true para que se ejecute de manera síncrona

	false para que se ejecute de manera asíncrona

browser: (*playwright.Browser) Se debe de pasar el navegador de tipo  que se utilizara en el scrapping

username: (string)

El nombre de usuario / identificador.

password: (string)
*/
func CreateAsyncScrapping(scrapFn func(*playwright.BrowserContext, string, string, ...string) (Result, error), context *playwright.BrowserContext, username string, password string) func(chan Result, chan error, ...string) {
	return func(cr chan Result, ce chan error, params ...string) {
		r, e := scrapFn(context, username, password)
		cr <- r
		ce <- e
	}
}

// Función para cerrar el scrapper
// playwright: (*playwright.Playwright) Se debe de pasar el propio playwright que se tiene que cerrar
// browser: (*playwright.Browser) Se tiene que pasar el propio navegador que se tiene que cerrar
// retorna error, nil en caso de no haber ningún error, caso contrario ocurrió un error cerrando el browser o playwright
func CloseScrapper(pw *playwright.Playwright, browser *playwright.Browser, context *playwright.BrowserContext) error {
	var err error
	if err := (*browser).Close(); err != nil {
		err = errors.New(fmt.Sprintf("could not close browser: %v", err))
	}
	if err := (*context).Close(); err != nil {
		err = errors.New(fmt.Sprintf("could not close context: %v", err))
	}
	if err := pw.Stop(); err != nil {
		err = errors.New(fmt.Sprintf("could not stop Playwright: %v", err))
	}
	return err
}

// Función para tomar una screenshot
// page *playwright.Page
// name string
// Se le pasa un pointer de la página que esté abierta en ese momento y el nombre de la screenshot
func screenshot(page *playwright.Page, name string) {
	(*page).Screenshot((playwright.PageScreenshotOptions{
		Path: playwright.String(name),
	}))
}

// Struct para guardar un resultado o un error
type Response struct {
	Res Result `json:"res"`
	Err error  `json:"err"`
}
