package main

import (
	"fmt"
)

func main() {

	arr := []int {34,32,56,64,89,90}

	for i,val := range arr {
		fmt.Printf("%d: %d\n", i, val)
	}



}