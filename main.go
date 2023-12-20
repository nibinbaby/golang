package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	greetings := "Hello there, friends !"

	// string method in std library to look for existance of strings in another string
	fmt.Println(strings.Contains(greetings, "e!"))
	// string method in std library to replace the occurance of the second parameter with the third
	fmt.Println(strings.ReplaceAll(greetings, "Hello", "hi"))

	fmt.Println(strings.ToUpper(greetings))    // convert to upper case
	fmt.Println(strings.Index(greetings, "e")) // find the index of the second parameter - first occurance
	fmt.Println(strings.Split(greetings, "e")) // split the string based on the second parameter

	fmt.Println(greetings)

	ages := []int{45, 33, 77, 12, 98, 6, 99, 4}

	sort.Ints(ages) // sort integers ascending order
	fmt.Println(ages)

	index := sort.SearchInts(ages, 33) // sort and search for the second parameter and return the index in the sorted array; if not found outputs the lastindex+1 value assuming the number will be coming soon
	fmt.Println(index)

	names := []string{"mario", "yovaikam", "reacher", "ashes", "indiana"}

	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "indiana"))
}
