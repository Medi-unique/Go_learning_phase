package main

import( "fmt")

func main() {
// for loop
for i:=0 ; i<10 ; i++ {
	fmt.Printf("Number %d\n ", i)
}

x :=0
for x <10{
	fmt.Println(x)
	x++
}
//   Fizz Buzz
for i :=1 ; i<=100 ;i++ {
	if i % 15 ==0{
		fmt.Println("FizzBuzz")
	} else if i % 5 ==0{
		fmt.Println("Buzz")
	} else if i%3==0{
		fmt.Println("Fizz")
	} else {
		println(i)
	}




}




}