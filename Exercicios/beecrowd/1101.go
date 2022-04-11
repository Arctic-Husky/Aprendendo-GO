package main

import(
  "fmt"
)

func main(){
  var num1 int
  var num2 int
  var menor, maior int
  var soma int

  for{
    fmt.Scanf("%d", &num1)
    fmt.Scanf("%d", &num2)

    if num1 <= 0 || num2 <= 0{
      return
    }

    if num1 <= num2{
      menor = num1
      maior = num2
    }else{
      menor = num2
      maior = num1
    }

    for i := menor; i <= maior; i++{
      fmt.Printf("%d ", i)
      soma += i
    }

    fmt.Printf("Sum=%d\n",soma)

    soma = 0
    menor = 0
    maior = 0
  }
}