package main

import(
  "fmt"
  "strings"
)

func main() {
  entrada := "()"
  var saida string

  entradaFiltrada := strings.Split(entrada, "\n")

  sliceEntradas := make([]string, 3)
  
  
}

func corrige(sliceEntradas []string, saida *string){
  for i,_ := range sliceEntradas{
    *saida += verifica(sliceEntradas[i])
  }
}

func verifica(entrada string) string{
  var ultimoParenteses byte
  
  for i,_ := range entrada {
    if entrada[i] == '(' || entrada[i] == ')'{
      ultimoParenteses = entrada[i]
    }
  }

  if ultimoParenteses == '('{
    return "incorrect\n"
  }

  return "correct\n"
}