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

	var student string
	fmt.Println("Enter the first name of the student you're looking for: ")
	_, err := fmt.Scanf("%s\n", &student)
	if err != nil {
		utils.Exit("Enter a valid name")
	}

	for careerName, careerCode := range careers {
		fmt.Println("\nLooking for students who completed a BA in ", careerName)
		go extract.GetStudents(careerCode, careerName, studentChan)
	}
	for i := 0; i < len(careers); i++ {
		students := <-studentChan
		allGraduates = append(allGraduates, students...)
	}

	fmt.Println("Search concluded: Printing Results")
	hasFound := false
	for _, graduate := range allGraduates {
		if strings.Contains(graduate.Name, strings.ToLower(student)) {
			fmt.Printf("%s -- %s -- %s\n", strings.ToTitle(graduate.Name), graduate.Career, graduate.Year)
			hasFound = true
		}
	}

	if !hasFound {
		fmt.Printf("No students whose name starts with %s were found.\n", student)
	}

}

//  todo: improve search to look only in first name
// interesting example: "Rosa"
// FOSADO GAYOU ROSA MARIA DEL CARMEN -- Business Administration -- 2011, we do want this result
// ROSADO MACHAIN JAVIER HUMBERTO  -- no
