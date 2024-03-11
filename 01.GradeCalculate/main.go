package main

import "fmt"


func sum(marks *[]int) int {
	var sum int
	for i := 0; i < len(*marks); i++ {
		sum += (*marks)[i]
	}
	return sum
}

func main() {

	// COnsole input in go Coopilot give me the code I want to accept image from user
	fmt.Println("Enter your name: ")
	var name string
	var numberOfSubjects int;
	fmt.Scanln(&name)
	fmt.Println("Enter the number of subjects: ")
	fmt.Scanln(&numberOfSubjects)
	var subjectName = make([]string, numberOfSubjects)
	var marks = make([]int, numberOfSubjects)
	for i := 0; i < numberOfSubjects; i++ {
		fmt.Println("Enter the subject name: ")
		fmt.Scanln(&subjectName[i])
		fmt.Println("Enter the marks: ")
		fmt.Scanln(&marks[i])
	}
	fmt.Println("Name: ", name)
	fmt.Println("Number of subjects: ", numberOfSubjects)
	for i := 0; i < numberOfSubjects; i++ {
		fmt.Println("Subject: ", subjectName[i], "Marks: ", marks[i])
	}
	var average = sum(&marks) / numberOfSubjects
	fmt.Println("Average: ", average)

}