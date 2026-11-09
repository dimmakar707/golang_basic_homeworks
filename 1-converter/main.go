package main

import "fmt"

func main() {
	const USDToEUR float64 = 0.8941
	const USDToRUB float64 = 82.14
	sumEUR := 100
	var sumRUB float64
	sumRUB = float64(sumEUR) / USDToEUR * USDToRUB
	fmt.Println(sumEUR, "евро равно", sumRUB)
}
