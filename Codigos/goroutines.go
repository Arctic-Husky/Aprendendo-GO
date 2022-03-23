package main

import (
  "fmt"
  "time"
)

// chan = channel, chan string = o canal carrega uma string
func simples(msg string, c chan string){
  for i := 0 ; i < 5 ; i++{
    c <- fmt.Sprintf(msg)
    //time.Sleep(time.Second)
  }
}

func main() {
  //go simples("Arctic")
  //go simples("Husky")
  //fmt.Printf("\tEsperando um pouco\n")
  c := make(chan string)
  go simples("Ping", c)
  for i := 0 ; i < 5 ; i++{
    fmt.Printf("Received: %q \n", <-c) // esperando receber mensagem pelo canal
    fmt.Printf("Pong\n");
    time.Sleep(1 * time.Second)
  }
  
  //time.Sleep(10 * time.Second)
  fmt.Printf("\n\tcabo\n")
}