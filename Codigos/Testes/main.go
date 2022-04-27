package main

import(
  "fmt"
  "errors"
)

const enPrefix = "Hello"
const ptPrefix = "Oi"

func main(){
  
}

func Hello(input string, lang string) (string, error){
  var prefix string

  switch lang{
    case "en": 
      prefix = enPrefix
    case "pt": 
      prefix = ptPrefix
    default: 
    return "", errors.New("Unknown Language")
  }
  
  return fmt.Sprintf("%s, %s!\n", prefix,input), nil
}