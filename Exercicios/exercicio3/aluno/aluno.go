package aluno

import (
  "fmt"
)

type Aluno struct{
  dataNascimento
  nome string
  matricula int
  disciplina
}

// Getters
func (a *Aluno) Nome()string {
  return a.nome
}

func (a *Aluno) Notas()[]int {
  return a.notas
}

type dataNascimento struct{
  dia int
  mes int
  ano int
}

type disciplina struct{
  notas []int
}

func (a *Aluno) MudarNota(nota int, notaAMudar int){
  a.notas[notaAMudar] = nota
}

func (a *Aluno) ExibirHistorico(){
  fmt.Printf("\n\tNotas: %v\n", a.notas);
}

func (a *Aluno) EditarNome(novoNome string){
  a.nome = novoNome
}

func (a *Aluno) LimparCadastro(){
  a.dia = 0
  a.mes = 0
  a.ano = 0
  a.notas = nil
  a.matricula = 0
  a = nil
}