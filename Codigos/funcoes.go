package main

import (
  "fmt"
  "log"
)

var rendimentoTinta float64

func CalculaTinta(largura float64, altura float64) (float64, error) {

  if largura <= 0 || altura <= 0 {
    return 0, fmt.Errorf("Largura e altura devem ser positivos maiores que 0.")
  }
  
  var area float64
  area = largura * altura
  
  /*if area < 0 {
    return 0, fmt.Errorf("Largura e altura devem ser positivos maiores que 0.")
  }*/
  
  litrosTinta := area/rendimentoTinta
  return litrosTinta, nil
}

func PrintaTinta(litrosTinta float64, erro error) {
  if erro != nil {
    log.Fatal(erro)
  }

  fmt.Printf("%.3f litros de tinta\n", litrosTinta)
}

func main() {
  rendimentoTinta = 12.1
  
  PrintaTinta(CalculaTinta(4.2, 5.11))
  PrintaTinta(CalculaTinta(2.1, 7.18))
}
