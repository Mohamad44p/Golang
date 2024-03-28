package main

import (
	"fmt"
)

func main() {
	age := 20
	name := "Mohammad"
	//printing
	fmt.Print("Hello,")
	fmt.Print(" World! \n")

	//Println
	fmt.Println("Hello, Mohammad!")
	fmt.Println("Hello, Omar!")
	fmt.Println("My age is", age, "and my name is", name)

	//Printf (formatted strings)

	//%v is for the value
	fmt.Printf("My age is %v and my name is %v\n", age, name)

	//%q is for the string
	fmt.Printf("My age is %q and my name is %q\n", age, name)

	//%T is for the type of the variable
	fmt.Printf("My age is %T and my name is %T\n", age, name)

	//%t is for the boolean
	isTrue := true
	fmt.Printf("The value of isTrue is %t\n", isTrue)

	//%d is for the decimal
	fmt.Printf("My age is %f\n", 255.55)

	//Sprintf (save formatted strings)
	var str = fmt.Sprintf("My age is %v and my name is %v\n", age, name)
	fmt.Println("The saved string is:", str)
}
