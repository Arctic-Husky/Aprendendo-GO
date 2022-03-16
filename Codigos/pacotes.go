package main

import (
  "fmt"
  //"log"
  "main/Codigos/Pacotes/data"
)

type evento struct{
  title string
  data.Date // campo anônimo
  // composição e agregação, evento tem comportamento de uma data agora. Não é herança exatamente
}

func main(){
  /*date := data.Date{}
  err := date.DateConstrutor(2022,03,16)
  if err != nil{
    log.Fatal(err)
  }*/

  //formatura := evento{"Dia da Formatura", date}
  formatura := evento{title: "Dia da Formatura"}
  formatura.DateConstrutor(2024,02,11)
  //fmt.Printf("%v\n", date.Year())
  fmt.Printf("%v\n", formatura)
}