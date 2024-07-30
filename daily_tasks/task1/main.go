package main

import (
	"fmt"
)

func average(subjects map[string]float64) float64 {
	sum := 0.0
	for _, grade := range subjects {
		sum += grade
	}
	return sum 
}

func main() {
	var totalSubjects int
	var fname string
	var lname string

	fmt.Println("Let's calculate your average grade!")
	fmt.Print("Enter your Full name: ")
	fmt.Scanln(&fname,&lname)
	fmt.Print("Enter the number of subjects: ")
	fmt.Scanln(&totalSubjects)

	subjects := make(map[string]float64)

	for i := 0; i < totalSubjects; i++ {
		var name string
		var grade float64
		fmt.Printf("Enter the name of subject %d: ", i+1)
		fmt.Scanln(&name)
		fmt.Printf("Enter your grade for %s: ", name)
		fmt.Scanln(&grade)

		if grade < 0 || grade > 100 {
			fmt.Println("Please enter a valid score between 0 and 100.")
			break
		} else {
			subjects[name] = grade
		}
	}
	if totalSubjects <= 0 {
		fmt.Println("No subjects entered. Cannot calculate average.")
		return
	}
	


	score := average(subjects)
	avg := score / float64(totalSubjects)

	row := fmt.Sprintf(" %-10s %-10s" ,fname, lname)
	fmt.Println(row)
	for key , val := range subjects {
		row := fmt.Sprintf("%-10s %-10f" ,key, val)
		fmt.Println(row)


	}
	r2  := fmt.Sprintf("Total %-10f ",score)
	fmt.Println(r2)
	r  := fmt.Sprintf("Average %-10f \n",avg)
	fmt.Println(r)
	

	switch {
	case score < 50.0:
		fmt.Printf("Your average score is %.2f. You should practice more and improve it.\n", score)
	case score <= 75.0:
		fmt.Printf("Your average score is %.2f. Good job! Keep practicing to improve even more.\n", score)
	default:
		fmt.Printf("Your average score is %.2f. You did a great job! Keep it up.\n", score)
	}
}
