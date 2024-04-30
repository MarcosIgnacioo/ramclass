package types

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
