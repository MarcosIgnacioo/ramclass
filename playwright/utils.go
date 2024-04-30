package pw

import (
	"fmt"

	"github.com/MarcosIgnacioo/utils"
)

type Assigment struct {
	ClassSubject string           `json:"class_subject"`
	Title        string           `json:"title"`
	Link         string           `json:"link"`
	Date         utils.DateFormat `json:"date"`
}

func NewAssigment(c string, t string, l string, d utils.DateFormat) Assigment {
	return Assigment{ClassSubject: c, Title: t, Link: l, Date: d}
}

type ScrappedInfo struct {
	Moodle    []interface{}
	ClassRoom []interface{}
}

type MoodleInfo []interface{}

func (si *ScrappedInfo) String() string {
	var assigments string
	for _, v := range si.Moodle {
		assigments += fmt.Sprintf("%v\n", v.(Assigment).Link)
	}
	return assigments
}

type LoginError struct {
	ErrorMessage string `json:"error_message"`
}

func NewLoginError(m string) *LoginError {
	return &LoginError{ErrorMessage: m}
}

func NewScrappedInfo(md []interface{}, cr []interface{}) *ScrappedInfo {
	return &ScrappedInfo{
		Moodle:    md,
		ClassRoom: cr,
	}
}
