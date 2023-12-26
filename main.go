package main

import "fmt"

func main() {

	for x := 0; x < 5; x++ {
		fmt.Println("value of x is :", x)
	}

	names := []string{"mario", "luigi", "yoshi", "peach"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for in loop ; using range keyword
	for index, value := range names {
		fmt.Printf("the position of name %v is %v \n", value, index)
	}

	// if you dont want to print out the index but only the value you can use _
	for _, value := range names {
		fmt.Printf("the value is %v \n", value)
		value = "new string" // if you alter the value in the loop it doesnt alter the original slice 'names'
	}

	fmt.Println(names)
}
