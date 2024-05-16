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
	ClassSubject string           `bson:"class_subject" json:"class_subject"`
	Title        string           `bson:"title" json:"title"`
	Link         string           `bson:"link" json:"link"`
	Date         utils.DateFormat `bson:"date" json:"date"`
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
	ClassRoom []interface{} `bson:"classroom" json:"classroom"`
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
	Moodle        []interface{} `bson:"moodle" json:"moodle"`
	ClassRoom     []interface{} `bson:"classroom" json:"classroom"`
	Kardex        []interface{} `bson:"kardex" json:"kardex"`
	CurricularMap []interface{} `bson:"curricular_map" json:"curricular_map"`
	Student       Result        `bson:"student" json:"student"`
	GPA           Result        `bson:"gpa" json:"gpa"`
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
func NewScrappedInfo(md []interface{}, cr []interface{}, kr []interface{}, cm []interface{}, st Result, gpa Result) *ScrappedInfo {
	return &ScrappedInfo{
		Moodle:        md,
		ClassRoom:     cr,
		Kardex:        kr,
		CurricularMap: cm,
		Student:       st,
		GPA:           gpa,
	}
}

type Error struct {
	ErrorMessage interface{} `bson:"error_message" json:"error_message"`
}

func (err *Error) GetResult() []interface{} {
	return []interface{}{err.ErrorMessage}
}

type LoginError struct {
	ErrorMessage interface{} `bson:"error_message" json:"error_message"`
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
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

// https://siia.uabcs.mx/siia2019/alumnos/credenciales.aspx?gr=alumno&op=photocrede
type StudentInfo struct {
	ControlNumber      int    `bson:"control_number" json:"control_number"`
	Name               string `bson:"name" json:"name"`
	InstitutionalEmail string `bson:"institutional_email" json:"institutional_email"`
	Campus             string `bson:"campus" json:"campus"`
	Career             string `bson:"career" json:"career"`
	Period             string `bson:"period" json:"period"`
	Semester           int    `bson:"semester" json:"semester"`
	Group              string `bson:"group" json:"group"`
	Turn               string `bson:"turn" json:"turn"`
	State              string `bson:"state" json:"state"`
}

func (s *StudentInfo) GetResult() []interface{} {
	return []interface{}{s.ControlNumber, s.Name, s.InstitutionalEmail, s.InstitutionalEmail, s.Campus, s.Career, s.Period, s.Semester, s.Semester, s.Group, s.Turn, s.State}
}

// Control:int, Number:string, Name:string, Email:string, Campus:string,
//
// Period:string, Semester:int, Group:string, Turn:string, State:string
func NewStudentInfo(cn int, n string, ie string, c string, ca string, p string, s int, g string, t string, st string) *StudentInfo {
	return &StudentInfo{ControlNumber: cn, Name: n, InstitutionalEmail: ie, Campus: c, Career: ca, Period: p, Semester: s, Group: g, Turn: t, State: st}
}

// https://siia.uabcs.mx/siia2019/alumnos/kardex.aspx?gr=alumno&op=kardex

type GPA struct {
	GPA int `gson:"gpa" json:"gpa"`
}

func NewGPA(gpa int) *GPA {
	return &GPA{GPA: gpa}
}

func (g *GPA) GetResult() []interface{} {
	return nil
}

func (g *GPA) GetGPA() int {
	return g.GPA
}

// Kardex types
type Kardex struct {
	GPA    int           `gson:"gpa" json:"gpa"`
	Kardex []interface{} `json:"kardex"`
}

func (k *Kardex) GetResult() []interface{} {
	return k.Kardex
}

func (k *Kardex) GetGPA() int {
	return k.GPA
}

func NewKardex(gpa int, k []interface{}) *Kardex {
	return &Kardex{GPA: gpa, Kardex: k}
}

type Subject struct {
	Semester int    `bson:"semester" json:"semester"`
	Subject  string `bson:"subject_name" json:"subject_name"`
	Group    string `bson:"group" json:"group"`
	Turn     string `bson:"turn" json:"turn"`
	Period   string `bson:"period" json:"period"`
	Grade    int    `bson:"grade" json:"grade"`
	State    string `bson:"state" json:"state"`
	Type     string `bson:"type" json:"type"`
	Teacher  string `bson:"teacher" json:"teacher"`
}

func NewSubject(s int, su string, g string, t string, p string, gr int, st string, ty string, te string) *Subject {
	return &Subject{Semester: s, Subject: su, Group: g, Turn: t, Period: p, Grade: gr, State: st, Type: ty, Teacher: te}
}

// Curricular map type

type CurricularSubject struct {
	Semester int         `bson:"semester" json:"semester"`
	Subject  string      `bson:"subject_name" json:"subject_name"`
	Period   interface{} `bson:"period" json:"period"`
	Grade    interface{} `bson:"grade" json:"grade"`
	State    interface{} `bson:"state" json:"state"`
	Credits  interface{} `bson:"credits" json:"credits"`
	Type     interface{} `bson:"type" json:"type"`
	Teacher  interface{} `bson:"teacher" json:"teacher"`
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
