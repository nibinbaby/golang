package main

import "fmt"

func main() {

	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	names := []string{"mario", "luigi", "bower", "yoshi", "peach"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue // go back up to the for loop next iteration
		}

		if index > 2 {
			fmt.Println("breaking at position", index)
			break // breaks out the loop completely
		}
		fmt.Printf("the value at pos %v is %v \n", index, value)
	}

}
