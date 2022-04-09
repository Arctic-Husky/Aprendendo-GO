package main

import(
  "fmt"
)

func main(){
  var A float64
  var B float64

  pesoA := 3.5
  pesoB := 7.5
  
  fmt.Scanf("%f", &A)
  fmt.Scanf("%f", &B)

  media := ((A * pesoA) + (B * pesoB))

  media = media / (pesoA + pesoB)
  
  fmt.Printf("MEDIA = %.5f\n", media)
}