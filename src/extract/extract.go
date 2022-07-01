package main

import (
	"findStudent/src/utils"
	"fmt"
	"github.com/gocolly/colly/v2"
)

func main() {
	careers := map[string]string{
		"finance":                 "001053",
		"economics":               "000038",
		"international relations": "000047",
		"math":                    "000169",
		"computer science":        "000009",
	}
	for careerName, careerCode := range careers {
		fmt.Println("\nLooking for students who completed a BA in ", careerName)
		students := GetStudents(careerCode)
		fmt.Println("First student: ", students[0])
	}

}

func GetStudents(careerCode string) []utils.Student {
	// Gets the name and date of graduation of all economics bachelors
	c := colly.NewCollector()
	var students []utils.Student

	c.OnHTML("table", func(e *colly.HTMLElement) {
		c.DetectCharset = true
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			newStudent := utils.Student{}
			row.ForEach("td", func(_ int, el *colly.HTMLElement) {
				switch el.Index {
				case 0:
					newStudent.Name = el.Text
				case 1:
					newStudent.Year = el.Text
					students = append(students, newStudent)
				}
			})

		})
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Request Status code: ", r.StatusCode)
	})

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL)
		r.ResponseCharacterEncoding = "ucs"
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	url := "http://escolar1.rhon.itam.mx/titulacion/titulados.asp?prog=" + careerCode
	err := c.Visit(url)
	if err != nil {
		return nil
	}

	fmt.Printf("Found %d graduated students\n", len(students))
	return students[1:] // first element contains the table headers
}
