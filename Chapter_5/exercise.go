package main

import "fmt"

var	x = []int {
	48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17,
}

func main() {

	min := x[0]

	for i := 0; i < len(x); i++ {

		if min > x[i] {

			min = x[i]

		}
		
	}

	fmt.Println("The minimum value within this array is:", min)

}