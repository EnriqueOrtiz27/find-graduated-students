package extract

import (
	"findStudent/src/utils"
	"fmt"
	"github.com/gocolly/colly/v2"
	"math"
	"strings"
)

func GetStudents(careerCode string, careerName string, channel chan []utils.Student) {
	// fmt.Println("\nLooking for students who completed a BA in ", careerName)
	c := colly.NewCollector()
	var students []utils.Student

	c.OnHTML("table", func(e *colly.HTMLElement) {
		c.DetectCharset = true
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			newStudent := utils.Student{}
			row.ForEach("td", func(_ int, el *colly.HTMLElement) {
				switch el.Index {
				case 0:
					newStudent.Name = strings.ToLower(el.Text)
				case 1:
					newStudent.Year = el.Text
					newStudent.Career = careerName
					students = append(students, newStudent)
				}
			})

		})
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Searching...")
	})

	c.OnRequest(func(r *colly.Request) {
		r.ResponseCharacterEncoding = "ucs" // the right format for these pages
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		utils.Exit(fmt.Sprintf("Request URL %s failed with error %v", r.Request.URL.Path, err))
	})

	url := "http://escolar1.rhon.itam.mx/titulacion/titulados.asp?prog=" + careerCode
	err := c.Visit(url)
	if err != nil {
		utils.Exit(fmt.Sprintf("Failed to visit page %s with error %v", url, err))
	}
	// data science has no graduates yet, and for the rest of the careers we need to start at index 1 to avoid
	// sending table headers into the channel
	index := math.Min(float64(1), float64(len(students)))
	channel <- students[int(index):]
}
