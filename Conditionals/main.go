package main

import (
	"fmt"
)

func main() {
	// if else
	x := 0
	y := 100
	col := "white"
	if x > y {
		fmt.Printf("%d is greater than %d", x, y)

	} else {
		fmt.Printf("%d is greater than %d", y, x)
	}
	// else if
	if x > y {
		fmt.Printf("%d is greater than %d", x, y)

	} else if x == y {
		fmt.Printf("%d is equall with  %d\n", y, x)
	} else {
		fmt.Printf("%d is greater than %d\n", y, x)
	}

	// switch

	switch col {
	case "white":
		fmt.Println("I like It")
	case "black":
		fmt.Println("I really liked It")
	default:
		fmt.Println("Nothing")

	}

}
