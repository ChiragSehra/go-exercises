package main

import (
  "fmt"
  "flag"
  "os"

)

func  main() {
  csvFileName := flag.String("csv","problems.csv","a csv file in the format of 'question, answer'")
  flag.Parse()

  file, err := os.Open(*csvFileName)
  if err!= nil {
    fmt.Println("Error:", err)
  }
  _ = file

}
