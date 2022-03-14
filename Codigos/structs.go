package main

import (
  "fmt"
)

type aluno struct {
    matricula int
    senha string
    status bool
    endereco
  }

type endereco struct {
  numero int
  rua string
}

var m int = 1

func main() {
  
  al := novoAluno("padrao")

  exibeAluno(al)

  expulsaAluno(&al)

  exibeAluno(al)
  
  /*var aluno struct {
    matricula int
    senha string
    status bool
  }*/

  /*type estudante struct {
    pessoa aluno
    tamanho float64
  }*/

  // var alunos [10]aluno

  /*for i,_ := range alunos{
    alunos[i] = novoAluno("padrao")
    exibeAluno(alunos[i])
  }*/

  // var a1 aluno

  // aluno.matricula = 4567
  // aluno.senha = "123abc"
  // aluno.status = true

  /*a1.matricula = 7531
  a1.senha = "4672csdfq"
  a1.status = true*/

  
  //fmt.Printf("%#v \n",alunos)
}

func exibeAluno(al aluno) {
  fmt.Printf("M: %d, S: %t \n",al.matricula, al.status)
  // fmt.Printf("Endereço: %s \n", al.endereco.rua)
  fmt.Printf("Endereço: %s \n", al.rua)
}

func novoAluno(senha string) aluno {
  alu := aluno{
    matricula:m,
    senha:senha,
    status:true,
  }

  alu.endereco = endereco {
    numero : 123,
    rua: "Rodosol",
  }

  m++

  return alu
}

func expulsaAluno(al *aluno) {
  al.status = false // em C precisaria da ->
}