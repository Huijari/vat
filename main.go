package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: ./vat [VAT NUMBER]")
		return
	}

	vat := os.Args[1]
	valid, err := CheckVat(vat[:2], vat[2:])
	if err != nil {
		fmt.Println(err)
		return
	}

	if valid {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}
}
