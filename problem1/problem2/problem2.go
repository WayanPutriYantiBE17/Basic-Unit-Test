package main

import "fmt"

func fibonacci(n int) []int {
	for _,value :=range n{}
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci((n - 2))
}

func main() {
	for i := 0; i <= 9; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println()
	for j := 0; j <= 4; j++ {
		fmt.Printf("%d ", fibonacci(j))
	}

}