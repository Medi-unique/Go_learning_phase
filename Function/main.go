package main

import (
	"fmt"
	"math"
)
func hello(s string){

	fmt.Println("Hello" ,s)
}
func bye(s string){
	fmt.Println("Bye" ,s)
}
func circleArea(r float64) float64 {
	return math.Pi * r*r 
}
func multi(a []string , f func(string) ){
	for _ , val := range a {
		f(val)


	}


}


func main() { 

	
	// hello("medi")
	// bye("Tofa")
	// multi([]string{"Medina","sofi","Haju"}, bye)
	a1 := circleArea(23)
	fmt.Printf( "%0.2f", a1)
	}

	
	
