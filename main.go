package main

import (
	"fmt"
	"math"
)

func main() {

	a := 10   // inferred type int
	b := 3.14 // inferred type float64
	//fmt.Println(a, b)

	c := float64(a)

	fmt.Println(c + b)

	fmt.Println(12345678910 + 20987654423)

	d := math.MaxFloat64 + 6665276555666677

	fmt.Println(d)

	g := 6
	h := 7

	if g > h {
		fmt.Println("76")
	} else if h > g+2 {
		fmt.Println("67")

	} else {
		fmt.Println("error")
	}

	/* for: loop
	   for i := 0; i < 100; i++ {
	       fmt.Println("Loop iteration:", i)
	   }*/

	numbers := [12345678]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Loop iteration:", i, " : ", numbers[i])
	}

}
