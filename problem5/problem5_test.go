package main

import "testing"

func TestPrintLuasPermukaan(t *testing.T) {
	t.Run("Test PrintLuasPermukaan",func(t *testing.T){
		var expected float64 = 602.88
		actual := PrintLuasPermukaan(4,20)
		if actual != expected{
			t.Error("Hasil tidak sesuai, actual:",actual,"expected:",expected)
		}
	}) 
	t.Run("Test PrintLuasPermukaan",func(t *testing.T){
		var expected float64 = 213.52
		actual := PrintLuasPermukaan(2,15)
		if actual != expected{
			t.Error("Hasil tidak sesuai, actual:",actual,"expected:",expected)
		}
	}) 
	t.Run("Test PrintLuasPermukaan",func(t *testing.T){
		var expected float64 = 226.08
		actual := PrintLuasPermukaan(4,5)
		if actual != expected{
			t.Error("Hasil tidak sesuai, actual:",actual,"expected:",expected)
		}
	}) 
	t.Run("Test PrintLuasPermukaan",func(t *testing.T){
		var expected float64 = 678.24
		actual := PrintLuasPermukaan(6,12)
		if actual != expected{
			t.Error("Hasil tidak sesuai, actual:",actual,"expected:",expected)
		}
	}) 
	t.Run("Test PrintLuasPermukaan",func(t *testing.T){
		var expected float64 = 282.6
		actual := PrintLuasPermukaan(5,4)
		if actual != expected{
			t.Error("Hasil tidak sesuai, actual:",actual,"expected:",expected)
		}
	}) 
	t.Run("Test PrintLuasPermukaan",func(t *testing.T){
		var expected float64 = 879.2
		actual := PrintLuasPermukaan(5,23)
		if actual != expected{
			t.Error("Hasil tidak sesuai, actual:",actual,"expected:",expected)
		}
	}) 
}
