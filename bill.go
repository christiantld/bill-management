package main

import (
	"fmt"
	"os"
)

type Bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill
func newBill(name string) Bill {
	bill := Bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return bill
}

func (b *Bill) format() string {
	formatedString := "Bill breakdown \n"
	var total float64 = 0

	// list items
	for key, value := range b.items {
		formatedString += fmt.Sprintf("%-25v ...$%0.2f \n", key+":", value)
		total += value
	}

	//tip
	formatedString += fmt.Sprintf("%-25v ...$%0.2f \n", "Tip:", b.tip)

	//total
	formatedString += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)

	return formatedString
}

func (b *Bill) updateTip(tip float64) {
	// same as (*b).tip = tip
	b.tip = tip

}

func (b *Bill) addItemToBill(item string, price float64) {
	//passing pointers avoid to create a copy of the map
	b.items[item] = price
}

func (b *Bill) saveBill() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		fmt.Println("Error creating file")
		panic(err)
	}
	fmt.Println("Bill was successfully saved")
}
