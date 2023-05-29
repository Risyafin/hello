package main

import "fmt"

func Penjumlahan(a int, b int) int {
	return a + b
}
func Pengurangan(a int, b int) int {
	return a - b
}
func Perkalian(a int, b int) int {
	return a * b
}

func pembagian(a int, b int) int {
	return a / b
}

func main() {
	var numberOne int
	var Operator string
	var numberTwo int
	var choice = "y"
	fmt.Println("---- Selamat datang di Kalkulator Sederhana ----")
	for choice == "y" {
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
		choice = "t"
		fmt.Println("Apakah anda ingin menggunakan kalkulator lagi ? (y/t)")
		fmt.Scan(&choice)

	}
}
