package main

import (
  "fmt"
)

func main() {
  num := 6
  
  //fmt.Printf("%v", &num)
  
  Duplicar(&num)
  fmt.Printf("%d\n", num)
}

func Duplicar(numero *int) {
  *numero *= 2 //*numero significa pra onde o ponteiro aponta, mexer no valor
}