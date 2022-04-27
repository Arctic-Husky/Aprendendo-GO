package main

import(
  "testing"
)

// "go test ./..."

// ciclo: consertar erros de compilação -> erro de teste, cabou

func TestHello(t *testing.T){
  // teste unitário
  t.Run("HelloBásico", func(t *testing.T){
  got,_ := Hello("World", "en")
  want := "Hello, World!\n"

  if got != want{
    t.Errorf("got %q, want %q\n", got, want)
  }
  })

  // teste unitário
  t.Run("HelloNome", func(t *testing.T){
  got,_ := Hello("Ricardo", "en")
  want := "Hello, Ricardo!\n"

  if got != want{
    t.Errorf("got %q, want %q\n", got, want)
  }
  })

  t.Run("HelloEmPortugues", func(t *testing.T){
  got,_ := Hello("Ricardo", "pt")
  // verificar se o erro é nulo
  want := "Oi, Ricardo!\n"

  if got != want{
    t.Errorf("got %q, want %q\n", got, want)
  }
  })

  t.Run("HelloIdiomaDesconhecido", func(t *testing.T){
  _,err := Hello("Ricardo", "portugues")
  if err == nil{
    t.Errorf("expected error, got none")
  } // verificar se retornou o erro correto
  })
}