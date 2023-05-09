package main

import "fmt"

func PrintLuas(alas, tinggi float64) float64 {
	// your code here
	luas := (alas*tinggi)/2
	return luas
}

func main() {
	var alas float64 = 20
	var tinggi float64 = 25

	fmt.Println(PrintLuas(alas, tinggi))

	var alas1 float64 = 8
	var tinggi1 float64 = 20

	fmt.Println(PrintLuas(alas1, tinggi1))

	var alas2 float64 = 3
	var tinggi2 float64 = 4

	fmt.Println(PrintLuas(alas2, tinggi2))

	var alas3 float64 = 4
	var tinggi3 float64 = 7

	fmt.Println(PrintLuas(alas3, tinggi3))

	var alas4 float64 = 6
	var tinggi4 float64 = 3

	fmt.Println(PrintLuas(alas4, tinggi4))
}
