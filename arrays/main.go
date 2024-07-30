package main
import(
	"fmt"
)
func greeting(s string) string {
	return "How are you"+s}
func prod(num1,num2 int) int {
	return num1 * num2}

func main(){
	arrr  := [4]string {"hello","world","world1","world2"}
	fmt.Println(greeting("Medi"))
	fmt.Println(prod(3,4))
	fmt.Println(arrr[2:])


}