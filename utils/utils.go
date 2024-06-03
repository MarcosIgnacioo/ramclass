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

var WeekDays = map[string]int{
	"domingo":   0,
	"lunes":     1,
	"martes":    2,
	"miércoles": 3,
	"jueves":    4,
	"viernes":   5,
	"sábado":    6,
	"Sunday":    0,
	"Monday":    1,
	"Tuesday":   2,
	"Wednesday": 3,
	"Thursday":  4,
	"Friday":    5,
	"Saturday":  6,
}

func GetDay(assigmentDay string) string {
	wDayCurr := WeekDays[time.Now().Weekday().String()]
	wDayAss := WeekDays[assigmentDay]
	var daysUntilAssigment int
	// Trabajo se entrega el jueves 4: dia: 7
	// Esto entra el sabado 6 dia:2
	// 0 1 2 3 4 5 6
	// d l m x j v s
	//             2
	// 3 4 5 6 7

	if wDayCurr > wDayAss {
		daysUntilAssigment = 7 - (wDayCurr - wDayAss)
	} else {
		daysUntilAssigment = wDayAss - wDayCurr
	}
	hours := fmt.Sprint(24*daysUntilAssigment, "h")
	duration, err := time.ParseDuration(hours)
	if err != nil {
		return fmt.Sprint("?")
	}
	return fmt.Sprint(time.Now().Add(duration).Day())
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
