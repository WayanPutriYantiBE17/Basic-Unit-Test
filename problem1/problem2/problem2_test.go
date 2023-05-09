package main

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	// Test case 1
	t.Run("Test 10", func(t *testing.T) {
		expected1 := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
		actual1 := make([]int, 10)
		for i := 0; i < 10; i++ {
			actual1[i] = fibonacci(i)
		}
		if !reflect.DeepEqual(actual1, expected1) {
			t.Errorf("Result was incorrect, got: %v, want: %v", actual1, expected1)
		}})
	
		t.Run("Test 3", func(t *testing.T) {
			fibonacci()
			expected1 := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
			actual1 := make([]int, 3)
			for i := 0; i < 10; i++ {
				actual1[i] = fibonacci(i)
			}
			if !reflect.DeepEqual(actual1, expected1) {
				t.Errorf("Result was incorrect, got: %v, want: %v", actual1, expected1)
			}
	})
}