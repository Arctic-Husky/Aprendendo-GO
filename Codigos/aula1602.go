package main

import (
  "fmt"
  //"time"
  "bufio"
  "os"
  "log"
  "strings"
  "strconv"
  
)

func main() {
  /*var now time.Time = time.Now()
  ano := now.Year()
  
  fmt.Println("Ano atual:", ano, ":)")*/

  fmt.Println("Entre com uma nota:  ")

  reader := bufio.NewReader(os.Stdin)
  //input, _ := reader.ReadString('\n') // o underline significa ignorar algum tipo de retorno ou algo que veio

  input, err := reader.ReadString('\n')

  

  if err != nil {
    log.Fatal(err)
  }/*else {
    fmt.Println(input)
  }*/

  input = strings.TrimSpace(input)

  nota, err := strconv.ParseFloat(input, 64)

  if err != nil {
    log.Fatal(err)
  }

  var status string
  
  if nota >= 7 {
    status = "Aprovado"
  }else if nota < 3 {
    status = "Reprovado"
  }else {
    status = "De final"
  }
  
  fmt.Println(status, nota)

  
}