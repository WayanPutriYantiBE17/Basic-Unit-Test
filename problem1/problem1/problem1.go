package main

import (
	"fmt"
	"math"
)
func angkaprima(n int) bool{
	if n < 2{
		return false
	}

	sqrtN := int(math.Sqrt(float64(n)))
	for i :=2; i<=sqrtN;i++{
		if n%i == 0{
			return false
		}
	}
	return true
}

func menemukanPrima(n,current int)int{
	if angkaprima(current){
		if n == 1{
			return current
		}
		return menemukanPrima(n-1,current+1)
	}
	return menemukanPrima(n,current+1)
}

func main(){
	n := 7 
	nthN := menemukanPrima(n,2)
	fmt.Printf("bilangan prima ke %d adalah %d",n,nthN)

	fmt.Println()
	n1 := 2
	nthN1 := menemukanPrima(n1,2)
	fmt.Printf("bilangan prima ke %d adalah %d",n1,nthN1)
}