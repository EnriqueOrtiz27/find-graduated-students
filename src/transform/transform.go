package main

import (
	"awesomeProject/src/extract"
	"awesomeProject/src/utils"
	"fmt"
	"strings"
)

func main() {
	x := extract.GetEconomists()
	result := parseNames(x)
	fmt.Println(result)
}

func parseNames(names string) []utils.Graduate {
	// receives a string containing the names and graduation dates of all bachelors and parses it
	graduate := utils.Graduate{Name: "Simple message", Year: "You're doing fine"}
	result := make([]utils.Graduate, 1)
	result[0] = graduate

	// try to split
	split := strings.Split(names, " ")
	fmt.Println("Split results: ", split[0])

	return result
}
