package main

import(
  "fmt"
)

func main(){
  var tamanho int
  var menorValor int
  var posicao int

  fmt.Scanf("%d", &tamanho)

  if tamanho <= 1 || tamanho >= 1000{
      return
  }
  
  vetorX := make([]int, tamanho)

  for i,_ := range vetorX{
    var num int
    fmt.Scanf("%d", &num)
    
    if i == 0{
      menorValor = num
    }

    if menorValor > num{
      menorValor = num
      posicao = i
    }

    vetorX[i] = num
  }

  fmt.Printf("Menor valor: %d\n", menorValor)
  fmt.Printf("Posicao: %d\n", posicao)
}