package main

import "testing"

func TestPrintNama(t *testing.T) {
	// your code here
	t.Run("Test kobar",func(t *testing.T){
		nama :="kobar"
		expected := "Hello kobar! Saya Golang bahasa yang sangat menyenangkan"
		actual := PrintNama(nama)
		if actual != expected{
			t.Error("Hasil tidak sesuai, actual:",actual,"expected:",expected)
		}
	}) 

	t.Run("Test ani",func(t *testing.T){
		nama1 :="ani"
		expected := "Hello ani! Saya Golang bahasa yang sangat menyenangkan"
		actual := PrintNama(nama1)
		if actual != expected{
			t.Error("Hasil tidak sesuai, actual:",actual,"expected:",expected)
		}
	}) 

	t.Run("Test santi",func(t *testing.T){
		nama2 :="santi"
		expected := "Hello santi! Saya Golang bahasa yang sangat menyenangkan"
		actual := PrintNama(nama2)
		if actual != expected{
			t.Error("Hasil tidak sesuai, actual:",actual,"expected:",expected)
		}
	}) 

	t.Run("Test dudu",func(t *testing.T){
		nama3 :="dudu"
		expected := "Hello dudu! Saya Golang bahasa yang sangat menyenangkan"
		actual := PrintNama(nama3)
		if actual != expected{
			t.Error("Hasil tidak sesuai, actual:",actual,"expected:",expected)
		}
	}) 

	t.Run("Test rani",func(t *testing.T){
		nama4 :="rani"
		expected := "Hello rani! Saya Golang bahasa yang sangat menyenangkan"
		actual := PrintNama(nama4)
		if actual != expected{
			t.Error("Hasil tidak sesuai, actual:",actual,"expected:",expected)
		}
	}) 
}
