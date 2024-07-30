package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main(){

	reader := bufio.NewReader(os.Stdin)


	fmt.Println("Enter the statment to be counted")
	statment,_ := reader.ReadString('\n')

	fmt.Println("Enter the statment to chceck weather palindrome or not")
	statment2,_ := reader.ReadString('\n')
	statment3 := strings.ToLower(statment2)
	
	resultForCount :=count(statment)
	fmt.Println(resultForCount)

	checkpalindrome := palindrome(statment3)
	fmt.Printf("the statment is %s \n",  checkpalindrome+".")




}