package main

import(
  "fmt"
  "strings"
)

func main(){
  var entrada string

  entradaData(&entrada)

  lista := divideData(entrada)

  printa(lista)
}

func entradaData(entr *string){
  fmt.Scanf("%s", entr)
}

func divideData(entr string)[]string{
  dividido := strings.Split(entr,"/")

  return dividido
}

func printa(entr []string){
  fmt.Printf("%s/%s/%s\n", entr[1], entr[0], entr[2])
  fmt.Printf("%s/%s/%s\n", entr[2], entr[1], entr[0])
  fmt.Printf("%s-%s-%s\n", entr[0], entr[1], entr[2])
}