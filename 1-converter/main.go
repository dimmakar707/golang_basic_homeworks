package main

import "fmt"

func main() {
	const USDToEUR float64 = 0.8941
	const USDToRUB float64 = 82.14
	const EURToRUB = USDToRUB / USDToEUR
	sumEUR := 100
	sumRUB := float64(sumEUR) * EURToRUB
	fmt.Println(sumEUR, "евро равно", sumRUB)
}
