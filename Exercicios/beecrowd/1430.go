package main

import(
  "fmt"
)

// compasso = soma das notas -> / notas /
// nota = tem uma duracao

const MAX int = 200
const MIN int = 3

var mapaNotas = map[byte]float64{
    'W': 1.0,
    'H': 1.0/2,
    'Q': 1.0/4,
    'E': 1.0/8,
    'S': 1.0/16,
    'T': 1.0/32,
    'X': 1.0/64,
  }

func main(){
  //notas := [7]byte{'W','H','Q','E','S','T','X'}
  //preencheDuracao(notas, mapaNotas)
  
  input,err := inputUsuario()
  if err != nil{
    return
  }
  
  lista := partirString(input)
  // for _,item := range lista{
  //   fmt.Printf("%c\n", item)
  // }
  qnt := 0
  comecou := false
  for _,uni := range lista{
    qnt = qnt + verificar(uni,mapaNotas)
    if uni != "" && !comecou{
      comecou = true
    }
    if uni == "" && comecou{
      fmt.Printf("%d\n",qnt)
      qnt = 0
    }
  }
  
}

func partirString(input string)[]string{
  var strings []string
  var str string
  for i,_ := range input{
    if input[i] != '/'{
      str += string(input[i])
    }else if input[i] == '/'{
      strings = append(strings, str)
      str = ""
    }
  }

  strings = append(strings, "")

  return strings
}

// func preencheDuracao(notas [7]byte,mapa map[byte]float64){
//   var anterior float64 = 1
//   for _,nota := range notas{
//     mapa[nota]++
//     mapa[nota] = anterior
//     anterior = anterior/2
//   }
// }

func inputUsuario() (string, error){
  var input string
  var inputCompleto string

  i := 0
  for{
    if i > MAX{
      return " ", fmt.Errorf("ERRO")
    }

    fmt.Scanf("%s", &input)

    inputCompleto += input

    if inputCompleto[len(inputCompleto)-1] == '*'{
      break
    }
    
    i++
  }

  if len(inputCompleto) > MAX || len(inputCompleto) < MIN{
    return " ", fmt.Errorf("ERRO")
  }

  return inputCompleto, nil
}

func verificar(input string, mapa map[byte]float64)int{
  soma := 0.0
  for i,_ := range input{
    soma += mapa[input[i]]
  }
  
  if modulo(soma-1) < 0.0001{
    return 1
  }

  return 0
}

func modulo(num float64)float64{
  if num < 0{
    num = num * -1.0;
  }

  return num
}