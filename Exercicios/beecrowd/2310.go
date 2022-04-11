package main

import(
  "fmt"
)

func main(){
  var numJogadores int
  var S, B, A, S1, B1, A1 float64
  var pS, pB, pA float64
  var somaS, somaB, somaA, somaS1, somaB1, somaA1 float64
  var nome string

  fmt.Scanf("%d\n", &numJogadores)

  for i := 0; i < numJogadores; i++{
    fmt.Scanf("%s", &nome)
    fmt.Scanf("%f %f %f", &S, &B, &A)
    fmt.Scanf("%f %f %f", &S1, &B1, &A1)

    somaS += S
    somaB += B
    somaA += A
    somaS1 += S1
    somaB1 += B1
    somaA1 += A1
  }
  pS = somaS1 / somaS * 100
  pB = somaB1 / somaB * 100
  pA = somaA1 / somaA * 100

  fmt.Printf("Pontos de Saque: %.2f %%.\nPontos de Bloqueio: %.2f %%.\nPontos de Ataque: %.2f %%.\n", pS, pB, pA)
}