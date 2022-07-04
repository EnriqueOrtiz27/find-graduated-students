package utils

import (
	"fmt"
	"os"
)

func GetCareers() map[string]string {
	// the page from which to extract the graduated students follows the following template
	// http://escolar1.rhon.itam.mx/titulacion/titulados.asp?prog={careerCode}
	return map[string]string{
		"Finance":                 "001053",
		"Economics":               "000038",
		"International Studies":   "000047",
		"Applied Mathematics":     "000169",
		"Computer Science":        "000009",
		"Actuarial Science":       "000031",
		"Data Science":            "001352",
		"Business Administration": "000032",
	}
}

func Exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
