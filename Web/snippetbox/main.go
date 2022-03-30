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
  mux.HandleFunc("/", home) // '/' funciona como um wildcard
  mux.HandleFunc("/snippet", showSnippet)
  mux.HandleFunc("/snippet/create", createSnippet)

  log.Printf("\n\tInicializando servidor na porta: \"4000\"\n")

  err := http.ListenAndServe(":4000", mux)

  log.Fatal(err)
}

func home(rw http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/"{
    http.NotFound(rw,r)
    return
  }
  rw.Write([]byte("Bem vindo ao SnippetBox"))
}

func showSnippet(rw http.ResponseWriter, r *http.Request) {
  rw.Write([]byte("Mostrando um snippet espicífico"))
}

func createSnippet(rw http.ResponseWriter, r *http.Request) {

  //para testar: curl -i -X POST http://localhost:4000
  if r.Method != "POST" {
    
    rw.Header().Set("Allow","POST")
    //rw.WriteHeader(405)
    http.Error(rw, "Metodo não permitido", http.StatusMethodNotAllowed)
    //rw.Write([]byte("Metodo não permitido"))
    return
  }
  rw.Write([]byte("Criando um novo snippet"))
}