package main

import "fmt"
import "strconv"

func main() {
  var angka1 string
  var angka2 string
  fmt.Print("Angka pertama :")
  fmt.Scanln(&angka1)
  fmt.Print("Angka kedua :")
  fmt.Scanln(&angka2)

  var number1 int
  var number2 int
  var err error
  number1, err = strconv.Atoi(angka1)
  number2, err = strconv.Atoi(angka2)

  if err == nil {
    hasiltambah := number1 + number2
    fmt.Println(fmt.Sprintf("Hasil %d + %d adalah %d", number1, number2, hasiltambah))

    hasilkurang := number1 - number2
    fmt.Println(fmt.Sprintf("Hasil %d - %d adalah %d", number1, number2, hasilkurang))

    hasilkali := number1 * number2
    fmt.Println(fmt.Sprintf("Hasil %d * %d adalah %d", number1, number2, hasilkali))

    hasilbagi := number1 / number2
    fmt.Println(fmt.Sprintf("Hasil %d / %d adalah %d", number1, number2, hasilbagi))
  } else {
    fmt.Println("error euy")
  }
}
