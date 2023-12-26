package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) { // return the initials capitalized
	if len(n) == 0 { // if n o argument as input retun empty strings
		return "", ""
	}
	s := strings.ToUpper(n)        // capitalize the string
	names := strings.Split(s, " ") // split the strings with space
	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1]) // append the first letter of each string of the slice 'names'
	}
	if len(initials) > 1 { // if initials length is more than 1 return the first and second elements by index
		return initials[0], initials[1]
	}

	return initials[0], "_" // if not greater than 1 return '_'

}

func main() {
	fn1, sn1 := getInitials("tifa lockharts")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitials("iron heart")
	fmt.Println(fn2, sn2)

	fn3, sn3 := getInitials("hustle")
	fmt.Println(fn3, sn3)

	fn4, sn4 := getInitials("")
	fmt.Println(fn4, sn4)
}
