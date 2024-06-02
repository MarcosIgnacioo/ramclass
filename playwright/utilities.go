package pw

import (
	"errors"
	"fmt"
	"log"

	"github.com/playwright-community/playwright-go"
)

var expect = playwright.NewPlaywrightAssertions(20000)
var loginSiiaExpect = playwright.NewPlaywrightAssertions(1000)
var await = playwright.NewPlaywrightAssertions(500)

/*
scrapFn: func(*playwright.Browser, string, string) (Result, error)

	La función que realizara el scrapping y que se quiera ejecutar, tiene que contar sus parametros deben de tener los siguientes tipos en el siguiente orden: *playwright.Browser, string, string y debe de retornar un Result y un error.

sync: (boolean)

Definirá si la función sera ejecutada de manera síncrona

	true para que se ejecute de manera síncrona

	false para que se ejecute de manera asíncrona

browser: (*playwright.Browser) Se debe de pasar el navegador de tipo  que se utilizara en el scrapping

username: (string)

	El nombre de usuario / identificador.

password: (string)

	La contraseña de la cuenta del usuario.
*/
func ChronoScrap(scrapFn func(*playwright.Browser, string, string) (Result, error), sync bool, browser *playwright.Browser, username string, password string) (Result, error, chan Result, chan error) {
	if !sync {
		cr := make(chan Result)
		ce := make(chan error)
		go func() {
			var r Result
			var e error
			r, e = scrapFn(browser, username, password)
			cr <- r
			ce <- e
		}()
		return nil, nil, cr, ce
	} else {
		res, err := scrapFn(browser, username, password)
		return res, err, nil, nil
	}
}

func CreateAsyncScrapping(scrapFn func(*playwright.BrowserContext, string, string, ...string) (Result, error), context *playwright.BrowserContext, username string, password string) func(chan Result, chan error, ...string) {
	return func(cr chan Result, ce chan error, params ...string) {
		r, e := scrapFn(context, username, password)
		cr <- r
		ce <- e
	}
}

// package main
//
// import "fmt"
//
//	func f(c chan func() (int, string)) {
//	    c <- (func() (int, string) { return 0, "s" })
//	}
//
//	func main() {
//	    c := make(chan func() (int, string))
//	    go f(c)
//	    y, z := (<-c)()
//	    fmt.Println(y)
//	    fmt.Println(z)
//	}
//
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

func screenshot(page *playwright.Page, name string) {
	(*page).Screenshot((playwright.PageScreenshotOptions{
		Path: playwright.String(name),
	}))
}

func logError(e *error) {
	if e != nil {
		log.Fatalf("error ocurred: %v", e)
	}
}

func closingError(logErr *error, browser *playwright.Browser, pw *playwright.Playwright) *LoginError {
	if logErr != nil {
		if err := (*browser).Close(); err != nil {
			log.Fatalf("could not close browser: %v", err)
		}
		if err := pw.Stop(); err != nil {
			log.Fatalf("could not stop Playwright: %v", err)
		}
		return NewLoginError((*logErr).Error())
	}
	return nil
}

type Response struct {
	Res Result `json:"res"`
	Err error  `json:"err"`
}

func Cronos(sync bool, browser *playwright.Browser, username string, password string, fn func(browser *playwright.Browser, username string, password string) (Result, error)) Response {
	if !sync {
		c := make(chan Response)
		go func() {
			var r Response
			r.Res, r.Err = fn(browser, username, password)
			c <- r
		}()
		return <-c
	} else {
		res, err := fn(browser, username, password)
		fmt.Println("this should go first")
		return Response{Res: res, Err: err}
	}
}

func Asynchronize(fn func(args ...interface{}) []interface{}) func(chan interface{}, interface{}) {
	return func(c chan interface{}, s interface{}) {
		res := fn(s)
		c <- res
	}
}
