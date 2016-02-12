package main

import "fmt"

var x [5]float64 //here we create array with 5 elements inside, 0 - 4.

func main() {

	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	//another way to create array is:
	y := [5] float64 {

		83,
		82,
		77,
		93,
		98,
		// 100, yes, we can directly comment a number out with '//', but remember to change the array length inside bracket. it's ture that array is hard to use inside Go, so usually we use slice instead.	
	} 

	var (
		total1 float64 = 0
		total2 float64 = 0
	)

	for i := 0; i < len(x); i++ {

		total1 += x[i]
		
	}

	fmt.Println(total1 / float64(len(x)))

	for _, value := range y { //the reason we are able to use range here is, the loop will go through the whole elements range of y, let's say if we try to loop the first 3, or 2 in the middle, then range is not applicable.

		total2 += value

	}

	fmt.Println(total2 / float64(len(y)))

	slice1 := []int{1, 2, 3}

	slice2 := append(slice1, 4, 5) //here we use append function to copy an existing array, and add more elements into this copied verion, make him a new one.
	
	fmt.Println(slice1, slice2)

	slice3 := []int{6, 7, 8}

	slice4 := make([]int, 2) //here we use make function to create a array that has no elements inside yet, but has 2 spot for elements hop-in later on.

	copy(slice4, slice3) //copy function us use as '(destination, source)', thus, here we copy slice3 elements into slice4, but since slice4 has only 2 spot, so the last element of slice3, 8, will not be copied over.

	fmt.Println(slice3, slice4)

	periodic := map[string]map[string]string { //map function comes, map[key]value, is the typical structure. here we use a map[key]value as a value of the former part map[key].
		
		"H" : map[string]string { //when we define map, we can either use 'make' and declear every map pair in each line, or use '{}' to contain every map pair in, use comma to seprate each pair up.
			
			"name" : "Hydrogen", //within a map pair, use colon between key and value, 'key : value', also, use comma after each map pair to seperate from the next one.
			"state" : "gas",
		
		},

		"He" : map[string]string {
			
			"name" : "Helium",
			"state" : "gas",
		
		},

		"Li" : map[string]string {
			
			"name" : "Lithium",
			"state" : "solid",
		
		},

		"Be" : map[string]string {
			
			"name" : "Beryllium",
			"state" : "solid",
		
		},

		"B" : map[string]string {
			
			"name" : "Boron",
			"state" : "solid",
		
		},

		"C" : map[string]string {
			
			"name" : "Carbon",
			"state" : "solid",
		
		},

		"N" : map[string]string {
			
			"name" : "Nitrogen",
			"state" : "gas",
		
		},

		"O" : map[string]string {
			
			"name" : "Oxygen",
			"state" : "gas",
		
		},

		"F" : map[string]string {
			
			"name" : "Fluorine",
			"state" : "gas",
		
		},

		"Ne" : map[string]string {
			
			"name" : "Neon",
			"state" : "gas",
		
		},

	}

	if element, ok := periodic["Li"]; ok { //here, use a condition within if, check if 'Li' is inside the periodic map we create. if true, then print element key and value accordingly, otherwise, do nothing.

		fmt.Println(element["name"], element["state"])
		
	}
}