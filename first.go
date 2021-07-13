package main

import "fmt"

func firstTask() {
	num := input()
	isEven(num)
}

func input() (num float64) {
	fmt.Println("Введитее число: ")
	fmt.Scanln(&num)
	return
}

func isEven(num float64) {
	if int(num)%2 == 0 {
		fmt.Println("Число чётное")
	} else {
		fmt.Println("Число не чётное")
	}
}
