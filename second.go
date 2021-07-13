package main

import "fmt"

func secondTask() {
	num := input()
	isDivideOn3(num)
}

func isDivideOn3(num float64) {
	if int(num)%3 == 0 {
		fmt.Println("Число подходит")
	} else {
		fmt.Println("Число не подходит")
	}
}
