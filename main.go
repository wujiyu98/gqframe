package main

import (
	"fmt"
	"strings"
)

func main() {
	// email.Send()
	price := 1234567890
	priceStr := fmt.Sprintf("%d", price)
	priceStr = priceStr[:-1]
	priceStr = priceStr[1:]
	priceStr = strings.Replace(priceStr, ".", ",", -1)
	fmt.Println(priceStr)

}
