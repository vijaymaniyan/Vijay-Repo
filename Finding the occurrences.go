package main

import "fmt"
//finding the occurrences of int value
func main() {

	
	occurences := []int{10, 20, 10, 30, 20, 40, 10}
//intialize the empty map 
	visit := make(map[int]int)
// iterate the occurrences
	for _, value := range occurences {

		visit[value] = visit[value] + 1
	}
	fmt.Println(visit)
}
