package main

import (
	"fmt"
	"math"
)

type rectangle struct { 
	width ,height float64
}
type circle struct {
	radius float64
}

type shape interface{
	area() float64
}

func (c circle) area() float64{
	return math.Pi*c.radius*c.radius
}
func (r rectangle) area() float64{

	return r.height * r.width
}

func main() {
  r1 := rectangle{23,7}
  c1 :=circle{3.0}
  shapes := []shape{c1 , r1}

  for _, shape := range shapes {
	fmt.Println(shape.area())
  }

  

}