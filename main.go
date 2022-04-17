package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)

	input, err := reader.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() Bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	bill := newBill(name)
	fmt.Println("Created the bill - ", bill.name)

	return bill
}

func promptOptions(bill Bill) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Choose option(a - add item, s - save bill, t - add tip, x - exit): ", reader)

	switch option {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Invalid price")
			promptOptions(bill)
		}

		bill.addItemToBill(name, p)
		fmt.Printf("Added %s to bill\n", name)
		promptOptions(bill)
	case "s":
		bill.saveBill()
	case "t":
		tip, _ := getInput("Tip: ", reader)
		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("Invalid tip")
			promptOptions(bill)
		}
		bill.updateTip(t)
		fmt.Printf("Tip: $%0.2f \n", t)
		promptOptions(bill)

	case "x":
		fmt.Println("Exiting")
	default:
		fmt.Println("Invalid option")
		promptOptions(bill)
	}
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}
