package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Word Frequency Count")
		fmt.Println("2. Palindrome Check")
		fmt.Println("3. Exit")

		fmt.Print("Enter your choice (1-3): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter a statement: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			wordCount := Frequency(input)
			fmt.Println("Word Frequency:", wordCount)
		case "2":
			fmt.Print("Enter a word or phrase: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if Palindrome(input) {
				fmt.Println("palindrome.")
			} else {
				fmt.Println("Not a palindrome.")
			}
		case "3":
			fmt.Println("Exiting")
			return
		default:
			fmt.Println("Please try again.")
		}
	}
}
