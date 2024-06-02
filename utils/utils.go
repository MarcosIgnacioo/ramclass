package utils

import (
	"fmt"
	"strings"
	"time"
)

var Months = map[string]string{
	"ene":       "enero",
	"feb":       "febrero",
	"mar":       "marzo",
	"abr":       "abril",
	"may":       "mayo",
	"jun":       "junio",
	"jul":       "julio",
	"ago":       "agosto",
	"sep":       "septiembre",
	"oct":       "octubre",
	"nov":       "noviembre",
	"dic":       "diciembre",
	"January":   "enero",
	"February":  "febrero",
	"March":     "marzo",
	"April":     "abril",
	"May":       "mayo",
	"June":      "junio",
	"July":      "julio",
	"August":    "agosto",
	"September": "septiembre",
	"October":   "octubre",
	"November":  "noviembre",
	"December":  "diciembre",
}

func Bytes_matrix_to_string(matrix [][]byte) string {
	var stringified string

	for _, v := range matrix {
		stringified += fmt.Sprint(string(v))
	}
	return stringified
}

func CreateDateClassroom(dateBytes [][]byte) *DateFormat {
	dateStringified := Bytes_matrix_to_string(dateBytes)
	return NewDateFormat(dateStringified)
}

func CreateDate(dateBytes [][]byte) *DateFormat {
	dateStringified := Bytes_matrix_to_string(dateBytes)
	return NewDateFormat(dateStringified)
}

type DateFormat struct {
	Day   string `json:"day"`
	Month string `json:"month"`
	Year  string `json:"year"`
	Hour  string `json:"hour"`
}

func NewDateFormat(dateString string) *DateFormat {
	date := strings.Split(dateString, " ")
	d := date[0]
	m := date[2]
	y := strings.Split(date[4], ",")[0]
	h := date[5]
	time.Now()
	return &DateFormat{Day: d, Month: m, Year: y, Hour: h}
}

func NewDateFormatClassroom(dateString string) *DateFormat {
	date := strings.Split(dateString, " ")
	d := date[0]
	m := date[2]
	y := strings.Split(date[4], ",")[0]
	h := date[5]
	return &DateFormat{Day: d, Month: m, Year: y, Hour: h}
}

func (df *DateFormat) String() string {
	return fmt.Sprintf(`%s de %s de %s Hora: %s`, df.Day, df.Month, df.Year, df.Hour)
}

type Result interface {
	GetResult()
}
