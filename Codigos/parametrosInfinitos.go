package main

import (
	"fmt"
	"math"
)

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	//fmt.Printf("Maior: %d", max(2, 4, 5, 7, 37, 5))
	fmt.Printf("Maior: %d", max(numeros...)) // passa um de cada vez no slice, funciona pra qualquer coisa
}

func max(nums ...int) int {
	max := int(math.Inf(-1))

	for _, n := range nums {
		if n > max {
			max = n
		}
	}

	return max
}
