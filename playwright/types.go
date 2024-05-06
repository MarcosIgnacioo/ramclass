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
	Moodle        []interface{} `json:"moodle"`
	ClassRoom     []interface{} `json:"classroom"`
	Kardex        []interface{} `json:"kardex"`
	CurricularMap []interface{} `json:"curricular_map"`
	Student       []interface{} `json:"student"`
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

// Cosas que cambiaria
// En vez de tener que usar interface{} lo que haria seria crear una propia interfaz generica que se llamase Info o algo asi
// Todos los structs de scrapping implementarian el metodo o metodos de dicha interfaz, y asi podemos usar todos en este struct por ejemplo en vez de ser interface{} serian sus tipos correspondientes, en realidad no se para que ocupo la interfaz generica, diria que para la funcion que retorna los datos pero creo que go lo hace por si mismo asi que pues deberia de experimentar con eso despues
func NewScrappedInfo(md []interface{}, cr []interface{}, kr []interface{}, cm []interface{}, st []interface{}) *ScrappedInfo {
	return &ScrappedInfo{
		Moodle:        md,
		ClassRoom:     cr,
		Kardex:        kr,
		CurricularMap: cm,
		Student:       st,
	}
}

type Error struct {
	ErrorMessage interface{} `json:"error_message"`
}

func (err *Error) GetResult() []interface{} {
	return []interface{}{err.ErrorMessage}
}

type LoginError struct {
	ErrorMessage interface{} `json:"error_message"`
}

func (le *LoginError) GetResult() []interface{} {
	return nil
}

func NewLoginError(m string) *LoginError {
	return &LoginError{ErrorMessage: m}
}

func NewError(m error) *LoginError {
	if m != nil {
		return &LoginError{ErrorMessage: m.Error()}
	}
	return &LoginError{ErrorMessage: nil}
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
	State              string `json:"state"`
}

func (s *StudentInfo) GetResult() []interface{} {
	return []interface{}{s.ControlNumber, s.Name, s.InstitutionalEmail, s.InstitutionalEmail, s.Campus, s.Career, s.Period, s.Semester, s.Semester, s.Group, s.Turn, s.State}
}

// Control number, name, email, campus, period, semester, group, turn
func NewStudentInfo(cn int, n string, ie string, c string, ca string, p string, s int, g string, t string, st string) *StudentInfo {
	return &StudentInfo{ControlNumber: cn, Name: n, InstitutionalEmail: ie, Campus: c, Career: ca, Period: p, Semester: s, Group: g, Turn: t, State: st}
}

// https://siia.uabcs.mx/siia2019/alumnos/kardex.aspx?gr=alumno&op=kardex

// Kardex types
type Kardex struct {
	Kardex []interface{} `json:"kardex"`
}

func (k *Kardex) GetResult() []interface{} {
	return k.Kardex
}

func NewKardex(k []interface{}) *Kardex {
	return &Kardex{Kardex: k}
}

type Subject struct {
	Semester int    `json:"semester"`
	Subject  string `json:"subject_name"`
	Group    string `json:"group"`
	Turn     string `json:"turn"`
	Period   string `json:"period"`
	Grade    int    `json:"grade"`
	State    string `json:"state"`
	Type     string `json:"type"`
	Teacher  string `json:"teacher"`
}

func NewSubject(s int, su string, g string, t string, p string, gr int, st string, ty string, te string) *Subject {
	return &Subject{Semester: s, Subject: su, Group: g, Turn: t, Period: p, Grade: gr, State: st, Type: ty, Teacher: te}
}

// Curricular map type

type CurricularSubject struct {
	Semester int         `json:"semester"`
	Subject  string      `json:"subject_name"`
	Period   interface{} `json:"period"`
	Grade    interface{} `json:"grade"`
	State    interface{} `json:"state"`
	Credits  interface{} `json:"credits"`
	Type     interface{} `json:"type"`
	Teacher  interface{} `json:"teacher"`
}

func (cs CurricularSubject) String() string {
	return fmt.Sprintf(`Semester: %v
		Subject: %v
		Period: %v
		Grade: %v
		State: %v
		Credits: %v
		Type: %v
		Teacher: %v
		`, cs.Semester, cs.Subject, cs.Period, cs.Grade, cs.State, cs.Credits, cs.Type, cs.Teacher)
}

func NewCurricularSubject(s int, su string, p interface{}, g interface{}, st interface{}, c interface{}, ty interface{}, te interface{}) *CurricularSubject {
	return &CurricularSubject{Semester: s, Subject: su, Period: p, Grade: g, State: st, Credits: c, Type: ty, Teacher: te}
}

type CurricularMap struct {
	CurricularMap []interface{}
}

func (c *CurricularMap) GetResult() []interface{} {
	return c.CurricularMap
}

// Void
func (c *CurricularMap) Print() {
	for _, v := range c.CurricularMap {
		fmt.Println(v.(CurricularSubject).String())
	}
}

func NewCurricularMap(c []interface{}) *CurricularMap {
	return &CurricularMap{CurricularMap: c}
}
