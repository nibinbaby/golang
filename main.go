package main

import "fmt"

//maps - allows us to save in key value format with keys and values of a particular type

func main() {

	//			 type	type
	//	  <-type key	value
	menu := map[string]float64{
		"soup":    99.4,
		"pie":     87.7,
		"salad":   88.33,
		"pudding": 33.44,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	//looping maps
	for k, v := range menu {
		fmt.Println(k, " - ", v)
	}

	//ints as key type

	phonebook := map[int]string{
		123:   "mario",
		423:   "luigi",
		1783:  "miss",
		12213: "bower",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[1783])

	//update the phonebook

	phonebook[423] = "superman"

	fmt.Println(phonebook)
}
