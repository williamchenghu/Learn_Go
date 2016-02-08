package	main //main is the main package we create here (thinking in a modularized way, then every package is a module that can be used by other ones).

import	"fmt" //here we import and able to use functions from another package called 'fmt', short of format

func main() { //create the function that will be executed within main package
	fmt.Println('Hello, World, this is my very first Go program!') //here is the real deal. a function 'Println', which will print a line on screen, inside fmt package, is called out here, to print a line for us, which is the Hello bla...bla...
}