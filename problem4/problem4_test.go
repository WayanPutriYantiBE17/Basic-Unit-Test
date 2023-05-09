package main

import "testing"

func TestPrintLuas(t *testing.T) {
	// your code here
	t.Run("Test PrintLuas",func(t *testing.T){
		var expected float64 = 250
		actual := PrintLuas(20,25)
		if actual != expected{
			t.Error("Hasil tidak sesuai, actual:",actual,"expected:",expected)
		}
	}) 
	t.Run("Test PrintLuas",func(t *testing.T){
		var expected float64 = 80
		actual := PrintLuas(8,20)
		if actual != expected{
			t.Error("Hasil tidak sesuai, actual:",actual,"expected:",expected)
		}
	}) 
	t.Run("Test PrintLuas",func(t *testing.T){
		var expected float64 = 6
		actual := PrintLuas(3,4)
		if actual != expected{
			t.Error("Hasil tidak sesuai, actual:",actual,"expected:",expected)
		}
	}) 
	t.Run("Test PrintLuas",func(t *testing.T){
		var expected float64 = 14
		actual := PrintLuas(4,7)
		if actual != expected{
			t.Error("Hasil tidak sesuai, actual:",actual,"expected:",expected)
		}
	}) 
	t.Run("Test PrintLuas",func(t *testing.T){
		var expected float64 = 9
		actual := PrintLuas(6,3)
		if actual != expected{
			t.Error("Hasil tidak sesuai, actual:",actual,"expected:",expected)
		}
	}) 
}
