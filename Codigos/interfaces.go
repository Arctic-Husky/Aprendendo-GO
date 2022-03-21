package main

import (
  "fmt"
  "math"
)

func main() {
  circ := circulo{raio:4.3}
  rect := retangulo{base:5.2, altura:12}

  medidas(circ)
  medidas(rect)
}

func medidas(fg formaGeometrica) {
  fmt.Printf("\n\tArea = %f\n\tPerimetro = %f\n", fg.area(), fg.perimetro());
}

type formaGeometrica interface {
  area() float64
  perimetro() float64
}

// Tipo circulo que implementa a Interface
type circulo struct {
  raio float64
}

func (c circulo) area() float64 {
  return math.Pi * c.raio * c.raio
}

func (c circulo) perimetro() float64 {
  return 2 * math.Pi * c.raio
}

// Tipo retangulo que implementa a Interface
type retangulo struct {
  base float64
  altura float64
}

func (r retangulo) area() float64 {
  return r.base * r.altura
}

func (r retangulo) perimetro() float64 {
  return (2 * r.altura) + (2 * r.base)
}