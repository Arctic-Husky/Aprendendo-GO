package main

import(
  "fmt"
  "strings"
  "os"
  "io"
)

const tamanhoPeq int = 1001
const tamanhoG int = 10001

func main() {
  //entrada := "a+(b*c)-2-a\n(a+b*(2-c)-2+a)*2\n(a*b-(2+c)\n2*(3-a))\n)3+b*(2-c)("
  var entrada string
  var saida string

  //var entradaChar [10000]byte

  entradaDados(&entrada)

  // for i,_ := range entradaChar{
  //   entrada += string(entradaChar[i])
  // }

  entradasDivididas := strings.Split(entrada, "\n")

  corrige(entradasDivididas, &saida)

  printa(saida)
  
}

func entradaDados(entrada *string){
  var char [tamanhoG]byte
  var charSmol [tamanhoPeq]byte
  //var ultimoChar byte

  if len(os.Args) < 2{
    return
  }

  tamanho, _ := os.Stat(os.Args[1])

  if tamanho.Size() == 0{
    return
  }

  arq,_ := os.Open(os.Args[1])

  buff := make([]byte, tamanho.Size())


  for{
    _, err := arq.Read(buff)
    if err != nil{
      if err == io.EOF{
        break
      }
      return
    }
  }

  k := 0

  for i,_ := range buff{
    if i < tamanhoG{
      charSmol[k] = buff[i]
      char[i] = charSmol[k]
      
      if charSmol[k] == '\n'{
        k = 0
      }

      k++
      
      if k >= tamanhoPeq{
        char[i+(k-1)] = '\n'
        k = 0
      }

      
    }

    if i >= tamanhoG{
      char[i-1] = '\n'
      break
    }
  }

  for m,_ := range char{
    *entrada += string(char[m])
  }
}
    
  

func printa(saida string){
  fmt.Printf("%s\n",saida)
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
      aberturas = append(aberturas, entrada[i]) // se for abertura salvar no vet
    }else if entrada[i] == ')'{ // se for fechadura verificar se ja tem abertura
      temParenteses = true
      achou = false
      for j,_ := range aberturas{
        if aberturas[j] == '('{
          achou = true
          break
        }
      }
      if !achou{ // se n tiver abertura pra fechadura, ta errado
        return "incorrect\n"
      }else{
        // se achou par de fechadura e abertura em aberturas, retirar abertura da """"pilha""""
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
  // se a """pilha""" nao estiver vazia, ta errado
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