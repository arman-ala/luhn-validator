package main

import (
	"fmt"

	luhn "github.com/arman-ala/luhn-validator/internalServices"
)

func main() {
	fmt.Println(luhn.Validate("555"))
}
