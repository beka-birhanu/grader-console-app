package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	gradesrvs "github.com/beka-birhanu/grader/app/grade"
	"github.com/beka-birhanu/grader/cmd/util"
	"github.com/beka-birhanu/grader/infrastructure"
)

// It prompts the user to enter their name and the number of subjects.
// For each subject, it asks for the subject name and grade.
// Finally, it calculates and displays the average grade.
func main() {
	reader := bufio.NewReader(os.Stdin)
	subjectStore := infrastructure.New()
	gradeService := gradesrvs.New(subjectStore)

	for {
		// Prompt for the student's name
		fmt.Print("Enter your name: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		// Prompt for the number of subjects, ensuring valid input
		var numSubjects int
		for {
			fmt.Print("Enter the number of subjects: ")
			numSubjectsStr, _ := reader.ReadString('\n')
			numSubjectsStr = strings.TrimSpace(numSubjectsStr)
			var err error
			numSubjects, err = strconv.Atoi(numSubjectsStr)
			if err != nil || numSubjects <= 0 {
				fmt.Println("Invalid number of subjects. Please enter a positive integer.")
			} else {
				break
			}
		}

		// Loop to get each subject's name and grade
		for i := 0; i < numSubjects; i++ {
			fmt.Printf("Enter the name of subject %d: ", i+1)
			subjectName, _ := reader.ReadString('\n')
			subjectName = strings.TrimSpace(subjectName)

			// Prompt for the grade, ensuring valid input
			var grade float64
			for {
				fmt.Printf("Enter the grade for %s: ", subjectName)
				gradeStr, _ := reader.ReadString('\n')
				gradeStr = strings.TrimSpace(gradeStr)
				var err error
				grade, err = strconv.ParseFloat(gradeStr, 64)
				if err != nil || grade < 0 || grade > 100 {
					fmt.Println("Invalid grade. Please enter a number between 0 and 100.")
				} else {
					break
				}
			}

			// Add the subject and grade to the grade service
			_, err := gradeService.Add(name, subjectName, float32(grade))
			if err != nil {
				fmt.Println(err.Error())
				i-- // To re-enter the current subject and grade
				continue
			}
		}

		// Calculate and display the average grade
		response := gradeService.Avarage(name)
		util.DisplayAvarageResponse(*response)

		// Prompt to continue or quit
		fmt.Print("Do you want to continue? (y/n): ")
		cont, _ := reader.ReadString('\n')
		cont = strings.TrimSpace(strings.ToLower(cont))
		if cont != "y" {
			fmt.Println("Goodbye!")
			break
		}
	}
}

