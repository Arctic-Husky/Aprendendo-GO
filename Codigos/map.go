package main

import (
	"fmt"
	"strings"
)

func main() {
	votos := "ricardo\ncaio\ngabriel\nmarcelo\nhenrique\nhenrique\nhenrique\ngabriel\nmarcelo"

	lista := strings.Split(votos, "\n")

	maps := make(map[string]int)

	for _, nome := range lista {
		maps[nome]++
	}

	fmt.Printf("%v\n", maps)
}
