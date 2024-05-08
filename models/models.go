package models

import pw "github.com/MarcosIgnacioo/playwright"

type Curricular struct {
	pw.CurricularMap
	Identifier string `bson:"identifier" json:"identifier"`
}

type Kardex struct {
	pw.Kardex
	Identifier string `bson:"identifier" json:"identifier"`
}

type Moodle struct {
	pw.Assigment
	ControlNumber string `bson:"control_number" json:"control_number"`
}

type ClassRoom struct {
	pw.ClassRoomInfo
	ControlNumber string `bson:"control_number" json:"control_number"`
}

type Activites struct {
	pw.Assigment
	ControlNumber string `bson:"control_number" json:"control_number"`
}

type Calendar struct {
	Date          Date   `bson:"date" json:"date"`
	Description   string `bson:"description" json:"description"`
	IsDone        bool   `bson:"is_done" json:"is_done"`
	ControlNumber string `bson:"control_number" json:"control_number"`
}

type Date struct {
	Day   string `bson:"day" json:"day"`
	Month string `bson:"month" json:"month"`
	Year  string `bson:"year" json:"year"`
}
