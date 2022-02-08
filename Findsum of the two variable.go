package main

import (
	"fmt"
)

func main() {
     
	fmt.Println(findsum([]int{10, 20, 30, 40, 50}, 50))
}
// find pairs with the given sum in an array(O(n))
func findsum(arr []int, target int) []int {
// create a empty map 
	sumvalue := make(map[int]int)
// iterate the sum value
	for i, val := range arr {

		_, ok := sumvalue[val]

		if ok {

			return []int{i, sumvalue[val]}

		}
		sumvalue[target-val] = i

	}

	return nil
}
