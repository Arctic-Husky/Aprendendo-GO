package main

import(
  "fmt"
  "strings"
)

func main() {
  var entrada string
  var saida string

  for{
    err := entradaDados(&entrada)
    if err != nil{
      return
    }

    if entrada == ""{
      return
    }

    entradasDivididas := strings.Split(entrada, "\n")

    corrige(entradasDivididas, &saida)

    printa(saida)

    saida = ""
    entrada = ""
  }
}

func entradaDados(entrada *string) (error){
  var charSmol [1000]byte

  for i,_ := range charSmol{
    _, err := fmt.Scanf("%c", &charSmol[i])
    if err != nil{
      return fmt.Errorf("ERRO")
    }
    if charSmol[0] == '\n'{
      return nil
    }
    if charSmol[i] == '\n'{
      break
    }

  }

  for l,_ := range charSmol{
    *entrada += string(charSmol[l])
  }

  //*entrada += input

  return nil
}

func printa(saida string){
  fmt.Printf("%s",saida)
}

func corrige(sliceEntradas []string, saida *string){
  var temp string
  for i,_ := range sliceEntradas{
    if(len(sliceEntradas[i]) >= 1){
      temp = verifica(sliceEntradas[i])
      if temp != "0"{
        *saida += temp
      }
    }
  }
}

func verifica(entrada string) string{
  var aberturas []byte

  top := 0
  achou := false
  temParenteses := false
  for i := 0; i < len(entrada); i++{
    if entrada[i] == '(' {
      temParenteses = true
      aberturas = append(aberturas, entrada[i])
    }else if entrada[i] == ')'{
      temParenteses = true
      achou = false
      for j,_ := range aberturas{
        if aberturas[j] == '('{
          achou = true
          break
        }
      }
      if !achou{
        return "incorrect\n"
      }else{
        temParenteses = true
        
        if aberturas[top] == '('{
          aberturas[top] = ' '
          top++
        }else if (top+1 >= len(aberturas)){
          aberturas[top] = ' '
        }else if aberturas[top] != '('{
          if aberturas[top+1] == '('{
            top++
            break
          }
          return "incorrect\n"
        }
      }
    }
  }
  for i,_ := range aberturas{
    if aberturas[i] == '('{
      return "incorrect\n"
    }
  }

  if temParenteses{
    return "correct\n"
  }
  
  return "0"
}