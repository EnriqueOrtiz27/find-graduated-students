package extract

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func main() {
	GetEconomists()
}

func GetEconomists() string {
	// Gets the name and date of graduation of all economics bachelors
	c := colly.NewCollector()
	economists := make([]string, 1)

	c.OnHTML("table", func(e *colly.HTMLElement) {
		c.DetectCharset = true
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			// for each line "tr" do amazing things
			row.ForEach("td", func(_ int, el *colly.HTMLElement) {
				switch el.Index {
				case 0:
					fmt.Println("Nombre", el.Text)
				case 1:
					fmt.Println("Año", el.Text)
				case 3:
					// Video and slides link
				}
			})

		})
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		r.ResponseCharacterEncoding = "ucs"
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit("http://escolar1.rhon.itam.mx/titulacion/titulados.asp?prog=000038")

	return economists[0] // first two elements are not related to the request
}
