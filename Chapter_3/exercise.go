package main

import "fmt"

var (
	fah float64
	cel float64
	foot float64
	meter float64
	)

func main() {

	fmt.Println("Please enter a degree value in Fahrenheit: ")

	fmt.Scanf("%f", &fah)

	cel	= (fah - 32) * 5 / 9

	fmt.Println(fah, "Fahrenheit is", cel, "in Celsius.")
	
	fmt.Println("Please enter a length value in foot: ")

	fmt.Scanf("%f", &foot)

	meter = foot * 0.3048

	fmt.Println(foot, "ft is", meter, "m.")
}