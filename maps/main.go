package main

import (
	"fmt"
)

func main(){
count := make(map[string]int)
count["M"]=1
count["e"]=1
count["d"]=2
count["i"]=3
fmt.Println(count)

check := map[string]string {"medi":"m", "count":"c", "sofi":"s"}
fmt.Println(check)



}