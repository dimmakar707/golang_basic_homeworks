package main

import "fmt"

const USDToEUR float64 = 0.8941
const USDToRUB float64 = 82.14
const EURToRUB = USDToRUB / USDToEUR

func main() {
	sum, currencyFrom, currencyTo := getUserInput()

	result := calculateResult(sum, currencyFrom, currencyTo)

	fmt.Printf("Результат: %.2f", result)
}

func getUserInput() (float64, string, string) {
	var sum float64
	var currencyFrom string
	var currencyTo string
	fmt.Print("Введите сумму: ")
	fmt.Scan(&sum)
	fmt.Print("Введите исходную валюту")
	fmt.Scan(&currencyFrom)
	fmt.Print("Введите целевую валюту")
	fmt.Scan(&currencyTo)
	return sum, currencyFrom, currencyTo
}

func calculateResult(sum float64, currencyFrom string, currencyTo string) float64 {
	var result float64
	return result
}
