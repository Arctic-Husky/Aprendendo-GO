package main

import(
  "fmt"
)

const tamanho int = 1100

func main() {
  var expressao [tamanho]byte
  
  i := 0
  for{
    if i < tamanho{
      fmt.Scanf(" %s", expressao)

      if(verifica(expressao)){
        fmt.Printf("correct\n")
      }else{
        fmt.Printf("incorrect\n")
      }
    }
  }
}

func verifica(expressao [tamanho]byte)bool{
  var qtdP int = 0

  if expressao[0] == ')'{
    return false
  }

  i := 0
  for{
    if (i >= cap(expressao)) && (qtdP < 0){
      break
    }

    if expressao[i] == '('{
      qtdP++
    }

    if expressao[i] == ')'{
      qtdP--
    }

    i++
  }

  if qtdP == 0{
    return true
  }else{
    return false
  }
}