package main

import "fmt"

func main() {
	
	//let's test some integers with calculation
	fmt.Println("1 + 1 =", 1 + 1)

	fmt.Println("2 - 2 =", 2 - 2)

	fmt.Println("3 * 3 =", 3 * 3)

	fmt.Println("4 / 4 =", 4 / 4)

	fmt.Println("3 % 2 =", 3 % 2)

	//let's test some string
	fmt.Println(len("Hello, World")) //len function calculates the length of this string "Hello, World", notice space counts as well

	fmt.Println("Hello, World"[1]) //this will shows the 2nd character of sting "Hello, World", notice the 1st is 0, not 1, so 'e' should be shown instead of 'H', however, character is shown by a byte, which is also an integer, so 101 should show up when execute this program

	fmt.Println("Hello, " + "World") //works almost like calculation on integers or floats. but those calculation doesn't make sense on strings, so instead, this will work as build up a sentence with multiple words.

	//let's test some booleans
	fmt.Println(true && true) //'&&' means AND

	fmt.Println(true && false)

	fmt.Println(true || true) //'||' means OR

	fmt.Println(true || false)

	fmt.Println(!true) //'!' means NOT
}