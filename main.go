package main

import "fmt"

func main() {

	// strings
	var nameOne string = "myname"
	var nameTwo = "newname"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "anothername"
	nameThree = "differentname"

	fmt.Println(nameOne, nameThree)

	nameFour := "fourthname"

	fmt.Println(nameFour)

	// numbers

	// ints

	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory

	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint = 12

	fmt.Println(numOne, numTwo, numThree)

	var scoreOne float32 = 45.23
	var scoreTwo float64 = 4345.45

	scoreThree := 343.3

	fmt.Println(scoreOne, scoreTwo, scoreThree)

}
