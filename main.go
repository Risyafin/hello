package main

import (
	"fmt"
	"strings"
)

func Penjumlahan(a float64, b float64) float64 {
	 c := a + b
	 return c
}
func Pengurangan(a float64, b float64) float64 {
	return a - b
}
func Perkalian(a float64, b float64) float64 {
	return a * b
}

func pembagian(a float64, b float64) float64 {
	return a / b
}

func main() {
	var numberOne, numberTwo float64
	var Operator string
	var bool bool
	var choice string
	fmt.Println("---- Selamat datang di Kalkulator Sederhana ----")
	for bool == false {
		fmt.Print("Masukka bilangan pertama : ")
		fmt.Scan(&numberOne)
		fmt.Print("Masukkan operator : ")
		fmt.Scan(&Operator)
		fmt.Print("Masukkan bilangan kedua : ")
		fmt.Scan(&numberTwo)
		if Operator == "+" {
			fmt.Println("Hasil : ", Penjumlahan(numberOne, numberTwo))
		} else if Operator == "-" {
			fmt.Println("Hasil : ", Pengurangan(numberOne, numberTwo))
		} else if Operator == "*" {
			fmt.Println("Hasil : ", Perkalian(numberOne, numberTwo))
		} else if Operator == "/" {
			fmt.Println("Hasil : ", pembagian(numberOne, numberTwo))
		} else {
			fmt.Println("Operator yang anda masukkan salah")
		}
		fmt.Println("Apakah anda ingin menggunakan kalkulator lagi ? (y/t)")
		fmt.Scan(&choice)
		if strings.ToUpper(choice) == "T" {
			break
		}

	}
	fmt.Println("Terima kasih")
}
