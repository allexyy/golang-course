package main

import "fmt"

const course = 76.5

func main() {
	var userSum float64

	fmt.Scanln(&userSum)
	fmt.Println(userSum * course)
}
