package utils

import (
	"fmt"
	"strings"
)

func bytes_matrix_to_string(matrix [][]byte) string {
	var stringified string

	for _, v := range matrix {
		stringified += fmt.Sprint(string(v))
	}
	return stringified
}

func CreateDate(dateBytes [][]byte) *DateFormat {
	dateStringified := bytes_matrix_to_string(dateBytes)
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
	return &DateFormat{Day: d, Month: m, Year: y, Hour: h}
}

func (df *DateFormat) String() string {
	return fmt.Sprintf(`%s de %s de %s Hora: %s`, df.Day, df.Month, df.Year, df.Hour)
}
