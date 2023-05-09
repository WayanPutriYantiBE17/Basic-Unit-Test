package main

import "fmt"

func PrintNama(nama string) string {
	// your code here
	var output string
	sapa := "Hello "
	deskripsi := "! Saya Golang bahasa yang sangat menyenangkan"
	
	output = sapa + nama + deskripsi
	return output
}

func main() {
	var nama string = "kobar"
	fmt.Println(PrintNama(nama))

	var nama1 string = "ani"
	fmt.Println(PrintNama(nama1))

	var nama2 string = "santi"
	fmt.Println(PrintNama(nama2))

	var nama3 string = "dudu"
	fmt.Println(PrintNama(nama3))

	var nama4 string = "rani"
	fmt.Println(PrintNama(nama4))
}
