package main

import (
  "fmt"
  "time"
  "bufio"
  "os"
  "strings"
  "strconv"
)

func main() {
  //vetorTime()
  //vetorString()
  //vetorFloat()
  //vetorSlice()
  sliceFloat()
}

func vetorTime() {
  var datas [10]time.Time

  datas[0] = time.Now()
  datas[1] = time.Unix(1257894000, 0)
  datas[2] = time.Now()

  fmt.Println(datas[0], datas[1])
}

func vetorString() {
  //var notasMusicais [3]string = [3]string{"do","re","mi"}
  notasMusicais := [7]string{"do","re","mi","fa","so","la","si"}

  //foreach
  /*for indice, valor := range VETOR {
    
  }*/

  //foreach
  for i, nota := range notasMusicais{
    fmt.Printf("[%d] -> %s\n",i ,nota)
  }

  //sem usar o indice, se quiser
  /*for _, nota := range notasMusicais{
    fmt.Printf("%s\n",nota)
  }*/
  
  /*for i:=0; i<len(notasMusicais); i++ {
    fmt.Printf("[%d] -> %s\n",i ,notasMusicais[i])
  }*/
}

func vetorFloat() {
  var pesos [5]float64
  var soma float64

  reader := bufio.NewReader(os.Stdin)

  for i:=0; i<len(pesos); i++ {
    input, _ := reader.ReadString('\n')

    input = strings.TrimSpace(input)

    peso, _ := strconv.ParseFloat(input, 64)

    pesos[i] = peso
  }
  
  for _, p := range pesos {
    soma += p
  }

  var media float64

  media = soma/float64(len(pesos))

  fmt.Printf("%f", media)
  
}

func vetorSlice() {
  //var mySlice []string
  
  mySlice := make([]string, 3) //vetor com 7 posicoes agora
  //mySlice := []string {"Ricardo", "Jose", "Mario"}
  mySlice[0] = "Ricardo"
  mySlice[1] = "Jose"
  mySlice[2] = "Mario"

  fmt.Printf("Slice: %v \n", mySlice)
}

func sliceFloat() {
  var pesos []float64
  var soma float64
  var peso float64

  for peso >= 0 {
    fmt.Printf("Digite um peso:")
    fmt.Scanf("%f", &peso)
    
    pesos = append(pesos, peso)
  }

  subSlice := pesos[ 0 : len(pesos)-1 ]

  for _, num := range subSlice {
    soma += num
  }

  media := soma/float64(len(subSlice))

  fmt.Printf("%f \n", subSlice)
  fmt.Printf("%.2f \n", media)
}