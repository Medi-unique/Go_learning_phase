package main

import ( "fmt" 
)

func update(s string) string {
   s = "updated"
   return s
}
func main(){
	s1 := "check"
	s2 := update(s1)
	fmt.Println(s2)


}