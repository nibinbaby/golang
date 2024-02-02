package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// format the bill

func (b *bill) format() string { // the first (..) defines the object to which this func is associated with
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ... $%v \n", k+":", v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-25v ... $%0.2f \n", "tip:", b.tip)
	total += b.tip
	// total
	fs += fmt.Sprintf("%-25v ... $%0.2f", "total:", total)

	return fs
}

// update tip
func (b *bill) updateTip(tip float64) { // to update the origingal item we use pointer
	//(*b).tip = tip // optional, go automatically does the dereferencing
	b.tip = tip
}

// add item
func (b bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save bill
func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 644)
	if err != nil {
		panic(err)
	}
}
