package main

import "testing"

func TestPrima(t *testing.T) {

	t.Run("Test Prima ke 7",func(t *testing.T){
		var expected bool = true
		actual := angkaprima(7)
		if actual != expected{
			t.Error("Hasil tidak sesuai, actual:",actual,"expected:",expected)
		}

	})

	t.Run("Test Prima ke 2",func(t *testing.T){
		var expected bool = true
		actual := angkaprima(2)
		if actual != expected{
			t.Error("Hasil tidak sesuai, actual:",actual,"expected:",expected)
		}

	})

}

func TestMenemukanPrima(t *testing.T) {

	t.Run("Test Prima ke 7",func(t *testing.T){
		var expected int = 17
		actual := menemukanPrima(7,2)
		if actual != expected{
			t.Error("Hasil tidak sesuai, actual:",actual,"expected:",expected)
		}

	})

	t.Run("Test Prima ke 2",func(t *testing.T){
		var expected int = 3
		actual := menemukanPrima(2,2)
		if actual != expected{
			t.Error("Hasil tidak sesuai, actual:",actual,"expected:",expected)
		}

	})

}