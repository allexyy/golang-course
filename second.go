package main

import (
	"fmt"
	"math"
)

func main() {

	sideA := 50
	sideB := 28

	sideC := math.Sqrt(math.Pow(float64(sideA), 2) + math.Pow(float64(sideB), 2))
	area := sideA * sideB / 2
	perimeter := float64(sideA) + float64(sideB) + sideC
	fmt.Println(area, perimeter, sideC)

}
