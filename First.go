package main

import (
	"fmt"
)

func main() {

	var price, delivery, sale uint64

	fmt.Print("Input a price: ")
	fmt.Scan(&price)

	fmt.Print("Input a price delivery: ")
	fmt.Scan(&delivery)

	fmt.Print("Input a sale: ")
	fmt.Scan(&sale)

	Total := (price + delivery - sale)

	fmt.Println("Total:", Total)

}
