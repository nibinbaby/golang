package main

import "fmt"

func main() {

	// arrays - the length is fixed
	// var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30} // shorthand

	names := [4]string{"mario", "peach", "bowser", "name"}
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices - the length can be varied (use arrays under the hood)

	var scores = []int{100, 40, 30}
	//append something - should be assigned to the same variable - it updates the variable with the new item
	scores = append(scores, 86)
	fmt.Println(scores, len(scores))

	//slice ranges
	rangeOne := names[1:3]  // the items from index 1 to index 3
	rangeTwo := names[2:]   // items from 2 to the last index
	rangeThree := names[:3] // items from first one to the third index

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)
}
