package main

import (
	"fmt"
)

func main() {
	//strings
	var nameOne string = "mohammad"
	var nameTwo = "Omar"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Emad"
	nameThree = "Ali"

	fmt.Println(nameOne, nameTwo, nameThree)

	namefour := "Ahmed"

	fmt.Println(nameOne, nameTwo, nameThree, namefour)

	//integers

	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	//bits & memory
	var numOne int8 = 25
	var numTwo int8 = 127
	var numThree uint16 = 255

	fmt.Println(numOne, numTwo, numThree)

	//floats

	var scoreOne float32 = 25.98
	var scoreTwo = 98.75
	scoreThree := 45.98

	fmt.Println(scoreOne, scoreTwo, scoreThree)
}
