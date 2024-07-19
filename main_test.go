package main

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		arr      []int
		expected []int
	}{
		{[]int{64, 34, 25, 12, 22, 11, 90}, []int{11, 12, 22, 25, 34, 64, 90}},
		{[]int{5, 1, 4, 2, 8}, []int{1, 2, 4, 5, 8}},
		{[]int{3, 0, 2, 5, -1, 4, 1}, []int{-1, 0, 1, 2, 3, 4, 5}},
		{[]int{}, []int{}}, 
		{[]int{1}, []int{1}}, 
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}}, 
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}}, 
	}

	for _, test := range tests {
		result := BubbleSort(test.arr)
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("BubbleSort(%v) = %v; want %v", test.arr, result, test.expected)
				break
			}
		}
	}
}
