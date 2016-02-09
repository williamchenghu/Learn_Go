package main

import "fmt"

func main() {

	fmt.Print("Enter a number: ")

	var input float64

	fmt.Scanf("%f", &input) //all we need to know now is, Scanf fills input number we enter, we'll get back to this later on.

	output := input * 2

	fmt.Print("This is the number you entered times 2:", output, ".")	
}