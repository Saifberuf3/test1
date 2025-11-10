package main

import "fmt"
import "math"

func main() {

	a := 10 // inferred type int
	b := 3.14 // inferred type float64
	//fmt.Println(a, b)
	
	c := float64(a)

	fmt.Println(c + b)
	
	text := "I AM A STRING ... "
	text2 := "I AM ANOTHER STRING"
	fmt.Println(text + text2)
	
	d := math.MaxFloat64
	fmt.Println(d)

	g := 6
	h := 7 

	if (g > h) {
		fmt.Println("76")
	} else if (h > g + 2) {
		fmt.Println("67")
	} else if (h > g - 2) {
		fmt.Println("OK")
	} else {
		fmt.Println("error")
	}


	    // for: loop
    for i := 0; i < 100; i++ {
        fmt.Println("Loop iteration:", i)
    }
}
