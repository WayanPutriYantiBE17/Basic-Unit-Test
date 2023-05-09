package main

import "fmt"

func PrintLuasPermukaan(radius, tinggi float64) float64 {
	// your code here

	var pi float64 = 3.14

	Lp := 2*pi*radius*(radius+tinggi)
	return Lp

}

func main() {
	var T float64 = 20
	var r float64 = 4

	fmt.Println(PrintLuasPermukaan(r, T))

	var T1 float64 = 15
	var r1 float64 = 2

	fmt.Println(PrintLuasPermukaan(r1, T1))

	var T2 float64 = 5
	var r2 float64 = 4

	fmt.Println(PrintLuasPermukaan(r2, T2))

	var T3 float64 = 12
	var r3 float64 = 6

	fmt.Println(PrintLuasPermukaan(r3, T3))

	var T4 float64 = 4
	var r4 float64 = 5

	fmt.Println(PrintLuasPermukaan(r4, T4))

	var T5 float64 = 23
	var r5 float64 = 5

	fmt.Println(PrintLuasPermukaan(r5, T5))
}
