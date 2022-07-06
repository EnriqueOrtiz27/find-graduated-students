package main

import (
	"findStudent/src/extract"
	"findStudent/src/utils"
	"fmt"
	"strings"
)

func main() {
	var allGraduates []utils.Student
	careers := utils.GetCareers()
	studentChan := make(chan []utils.Student)

	var firstName, lastName string
	fmt.Println("Do not use accents please!")
	firstName = utils.ReadUserInput("Enter the first name of the student you're looking for: ")
	lastName = utils.ReadUserInput("Enter the last name of the student you're looking for: ")
	fmt.Println()

	for careerName, careerCode := range careers {
		go extract.GetStudents(careerCode, careerName, studentChan)
	}
	for i := 0; i < len(careers); i++ {
		students := <-studentChan
		allGraduates = append(allGraduates, students...)
	}

	fmt.Printf("\nThese are the names most similar to \"%s %s\":\n\n", firstName, lastName)
	hasFound := false
	for _, graduate := range allGraduates {
		if utils.MatchNames(strings.ToLower(graduate.Name), firstName, lastName) {
			// strings.Title is deprecated, but it's harmless in this context
			fmt.Printf("%s: Graduated from %s in %s\n", strings.Title(graduate.Name), graduate.Career, graduate.Year)
			hasFound = true
		}
	}

	if !hasFound {
		fmt.Printf("No students whose name starts with %s were found.\n", firstName)
	}

}
