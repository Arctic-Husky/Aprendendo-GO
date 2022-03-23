package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  //"time"
  //"log"
)

func size(url string, c chan int){
  fmt.Printf("\tPegando o site: %s\n", url)

  response, _ := http.Get(url)

  // defer: executa no final da função
  defer response.Body.Close()

  body, _ := ioutil.ReadAll(response.Body)

  fmt.Printf("\tTamanho: %d de %s\n", len(body), url)
  c <- len(body)
}

func main() {
  c := make(chan int)
  
  go size("https://replit.com/@ArcticHusky/Aprendendo-GO#Codigos/web.go", c)
  go size("https://www.level5.co.jp/", c)
  go size("https://pt.wikipedia.org/wiki/Linus_Torvalds", c)
  go size("https://discord.com", c)
  go size("https://arctic-husky.github.io/", c)
  go size("https://classic.minecraft.net/", c)

  var total int
  total += <-c
  total += <-c
  total += <-c
  total += <-c
  total += <-c
  total += <-c

  fmt.Printf("Tamanho Total = %d\n",total)
}