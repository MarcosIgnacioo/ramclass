package pw

import (
	"fmt"

	"github.com/MarcosIgnacioo/utils"
)

// Assigments Types

type Result interface {
	GetResult() []interface{}
}

type Assigment struct {
	ClassSubject string           `json:"class_subject"`
	Title        string           `json:"title"`
	Link         string           `json:"link"`
	Date         utils.DateFormat `json:"date"`
}

func NewAssigment(c string, t string, l string, d utils.DateFormat) Assigment {
	return Assigment{ClassSubject: c, Title: t, Link: l, Date: d}
}

// Moodle Types

type MoodleInfo struct {
	Moodle []interface{} `json:"moodle"`
}

func (mi *MoodleInfo) GetResult() []interface{} {
	return mi.Moodle
}

func (mi *MoodleInfo) String() string {
	var assigments string
	for _, v := range mi.Moodle {
		assigments += fmt.Sprintf("%v\n", v.(Assigment).Link)
	}
	return assigments
}

func NewMoodleInfo(mi []interface{}) *MoodleInfo {
	return &MoodleInfo{Moodle: mi}
}

// ClassRoom Types

type ClassRoomInfo struct {
	ClassRoom []interface{} `json:"classroom"`
}

func (ci *ClassRoomInfo) GetResult() []interface{} {
	return ci.ClassRoom
}

func (ci *ClassRoomInfo) String() string {
	var assigments string
	for _, v := range ci.ClassRoom {
		assigments += fmt.Sprintf("%v\n", v.(Assigment).Link)
	}
	return assigments
}

func NewClassRoomInfo(cr []interface{}) *ClassRoomInfo {
	return &ClassRoomInfo{ClassRoom: cr}
}

// ScrappedInfo types

type ScrappedInfo struct {
	Moodle    []interface{} `json:"moodle"`
	ClassRoom []interface{} `json:"classroom"`
}

func (si *ScrappedInfo) GetResult() []interface{} {
	return append(si.ClassRoom, si.Moodle...)
}

func (si *ScrappedInfo) String() string {
	var assigments string
	for _, v := range si.Moodle {
		assigments += fmt.Sprintf("%v\n", v.(Assigment).Link)
	}
	return assigments
}

func NewScrappedInfo(md []interface{}, cr []interface{}) *ScrappedInfo {
	return &ScrappedInfo{
		Moodle:    md,
		ClassRoom: cr,
	}
}

type LoginError struct {
	ErrorMessage string `json:"error_message"`
}

func (le *LoginError) GetResult() {}

func NewLoginError(m string) *LoginError {
	return &LoginError{ErrorMessage: m}
}

// User Types

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// https://siia.uabcs.mx/siia2019/alumnos/credenciales.aspx?gr=alumno&op=photocrede
type StudentInfo struct {
	ControlNumber      int    `json:"control_number"`
	Name               string `json:"name"`
	InstitutionalEmail string `json:"institutional_email"`
	Campus             string `json:"campus"`
	Career             string `json:"career"`
	Period             string `json:"period"`
	Semester           int    `json:"semester"`
	Group              string `json:"group"`
	Turn               string `json:"turn"`
}

// https://siia.uabcs.mx/siia2019/alumnos/kardex.aspx?gr=alumno&op=kardex

type Kardex struct{}
