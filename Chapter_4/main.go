package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {

		fmt.Println(i)

	}

	for i := 1; i <= 10; i++ {

		if i % 2 == 0 {

			fmt.Println(i, "is even.")

		} else {

			fmt.Println(i, "is odd.")

		}
		
	}

	for i := 0; i <= 6; i++ {

		if i == 0 {

			fmt.Println(i, "is Zero.")

		} else if i == 1 {

			fmt.Println(i, "is One.")

		} else if i == 2 {

			fmt.Println(i, "is Two.")
			
		} else if i == 3 {

			fmt.Println(i, "is Three.")
			
		} else if i == 4 {
			
			fmt.Println(i, "is Four.")

		} else if i == 5 {

			fmt.Println(i, "is Five.")
			
		} else {

			fmt.Println(i, "is unknown number, because I can only count till five in English.")
		}
		
	}

	for i := 0; i <= 6; i++ {

		switch i {

			case 0: fmt.Println(i, "is Zero in ENG.")
			case 1: fmt.Println(i, "is One in ENG.")
			case 2: fmt.Println(i, "is Two in ENG.")
			case 3: fmt.Println(i, "is Three in ENG.")
			case 4: fmt.Println(i, "is Four in ENG.")
			case 5: fmt.Println(i, "is Five in ENG.")
			case 6: fmt.Println(i, "is... Six I guess?")
			default: fmt.Println("Unknown Number")
		}
		
	}
	
}