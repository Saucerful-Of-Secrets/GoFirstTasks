package main

import (
	"fmt"
)

func main() {

	var price, delivery, sale int

	fmt.Print("Input a price: ")
	fmt.Scan(&price)

	fmt.Print("Input a price delivery: ")
	fmt.Scan(&delivery)

	fmt.Print("Input a sale: ")
	fmt.Scan(&sale)

	sum := (price + delivery - sale)

	fmt.Println(sum)

}
