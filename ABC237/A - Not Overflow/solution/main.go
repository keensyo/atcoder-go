package main

import (
	"fmt"
	"math"
)

func main() {
	var N int64
	fmt.Scan(&N)
	floatN := float64(N)

	powMinN := math.Pow(-2.0, 31.0)
	powMaxN := math.Pow(2.0, 31.0)

	if (powMinN < floatN) && (floatN < powMaxN) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
