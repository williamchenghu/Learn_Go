package main

import "fmt"

var universal = "We call him universal is because he is universal!" //tart from line 42, see more details.

func main() {

	//let's introduce variables, so var is like a storage, we can put things inside can call it out later.
	var x string = "Hello, World"
	
	//Or, we can break above definition down to 2 lines
	var y string
	y = "Hello, World, Again!"

	fmt.Println(x)

	fmt.Println(y)

	//reason of calling it a variable is because the value assigned to it, can change time to time
	y = y + "[Same var y has changed now]"

	fmt.Println(y)

	//in fact, 'y = y + XXX' can be directly write as 'y += XXX'. this will do the same job
	x += "[var x has changed now also]"

	fmt.Println(x)

	//additionally, a definition followed by a variable is not essential, for instance, definition 'string' can be earsed after 'var x' because Go language can decided which definition it should assign to variable x due to the value user assign to x.
	var z = "Let's ingore giving z a definition!" //variable z will be recognized as string

	alpha := "This works either." //a ':=' will be a short cut as well, aplha will be recognized as string as well

	beta := 1024 //also, beta will be recognized as integer in this case

	fmt.Println(z, alpha, beta)

	//a constant means sth that would not change later on
	const gama = "Gama is a constant that cannot be changed later on!"

	fmt.Println(gama)

	//also, a shorthand way of giving multiple defines at a time is shown below
	var (
		a = 7
		b = 8
		c = 9
		)
	//this also works for constants
	fmt.Println("7 + 8 + 9 =", a + b + c)
	
	slave()

}

/*
now we talk about Scope, we defined x, y, z, alpha, beta, all these, inside the function 'main'.

If we create another function here, let's call him function slave, when we want to print a line

with x, y, z, alpha, beta, anyone of them, the program will return an error, which tells the var

we are calling is undefined. this is due to we gave defines inside the function main. those ones

can only be used within main fuction. To make a define universal, at least universal inside this package,

we shall define that before we creates functions, just like what we did to var 'universal', now

we should be able to call him up in any function we create.
*/

func slave() {

	fmt.Println(universal)
	
}