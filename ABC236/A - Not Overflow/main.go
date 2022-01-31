package main

import (
	"fmt"
	"math"
)

func main() {
	var N int64
	fmt.Scan(&N)
	floatN := float64(N)
	logN := math.Log2(floatN)

	if logN < 31.0 {
		fmt.Println("Yes")
	} else if math.IsNaN(logN) {
		reN := -floatN
		logN = math.Log2(reN)
		if logN > 31.0 {
			fmt.Println("No")
		} else {
			fmt.Println("Yes")
		}
	} else {
		fmt.Println("No")
	}
}
