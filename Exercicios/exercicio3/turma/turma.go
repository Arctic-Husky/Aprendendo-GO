package turma

import (
  "fmt"
  "main/Exercicios/exercicio3/aluno"
)

type Turma struct {
  codigo int
  listaAlunos []aluno.Aluno
  maximoAlunos int
}

func (t *Turma) AdicionarAluno(al aluno.Aluno){
  
}

func (t *Turma) RemoverAluno(pos int){
  
}

func (t *Turma) AdicionarNota(pos int, nota int){
  
}

func (t *Turma) ExibirResumo(){
  for i,_ : range t.listaAlunos{
    fmt.Printf("\n\tAluno: %s,\n\tNotas: %v\n", t.listaAlunos[i].Nome(), t.listaAlunos[i].Notas())
  }
}