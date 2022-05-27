package main

import (
	"fmt"
)

func plus_500(atm *int) int {
	*atm = *atm + 500
	return *atm
}

func main() {
	atm := 3
	fmt.Println("atm=", atm)
	atm_plus_500 := plus_500(&atm)
	fmt.Println("atm=", atm)
	fmt.Println("atm_plus_500=", atm_plus_500)
}
