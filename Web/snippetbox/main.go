package main

import (
  "net/http"
  "log"
)

//router www.site.com/coisa1/coisa2
//handlers tratar de requisições
//servidor

func main() {
  //multiplexador em concorrência, o escutador
  mux := http.NewServeMux()
  //uma rota, chamar "home" a cada /
  mux.HandleFunc("/", home)

  log.Printf("\n\tInicializando servidor na porta: \"4000\"\n")

  err := http.ListenAndServe(":4000", mux)

  log.Fatal(err)
}

func home(rw http.ResponseWriter, r *http.Request) {
  rw.Write([]byte("Bem vindo ao SnippetBox"))
}