package main

import "fmt"

// create a fibonacci series
func main() {
	
	a := 0
	b := 1
	for i := 0; i < 10; i++ {

		a, b = b, a+b

		fmt.Println(a)

	}

}
