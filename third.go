package main

import "fmt"

func main() {

	var debit float64
	var profitPercent float64

	fmt.Println("Вклад")
	fmt.Scanln(&debit)
	fmt.Println("Процент по вкладу")
	fmt.Scanln(&profitPercent)

	result := debit + debit*profitPercent/100*5
	fmt.Printf("Profit for 5 years: %.2f", result)
}
