package main

import (
	"fmt"
)

func Permute(nums []int) [][]int {
	var result [][]int
	if len(nums) == 0 {
		result = make([][]int, 0)
	} else if len(nums) < 2 {
		result = [][]int{nums}
	} else 	if len(nums) == 2 {
		result = [][]int{nums, []int{nums[1], nums[0]}}
	} else {
		for i := range nums {
			x := nums[i]
			startBuffer := nums[:i]
			endBuffer := nums[i+1:]
			start := make([]int, len(startBuffer))
			copy(start, startBuffer)
			
			end := make([]int, len(endBuffer))
			copy(end, endBuffer)

			y := append(start, end...)
			m := Permute(y)
			for j := range m {
				m[j] = append([]int{x}, m[j]...)
			}
			result = append(result, m...)
		}
	}

	return result
}

func Permute2(num []int) [][]int {
	possiblePosition := ComputePossiblePosition(len(num))
	m := make([][]int, possiblePosition)
	fmt.Println("possiblePosition", possiblePosition)
	// Allocation
	for i, _ := range m {
		n := make([]int, len(num))
		m[i] = n
	}
	fmt.Println("Matrix After init:")
	PrintMatrix(m)
	for i, x := range num {
		// for each row
		idx := i
		for k, _ := range m {
			idx %= len(num)

			m[k][idx] = x

			idx++
		}

	}

	fmt.Println("Matrix After filling:")
	PrintMatrix(m)

	return m
}

func ComputePossiblePosition(x int) int {
	if x <= 0 {
		return 0
	} else if x == 1 {
		return 1
	}
	return x * ComputePossiblePosition(x-1)
}

func PrintMatrix(m [][]int) {
	for i, _ := range m {
		fmt.Printf("%+v\n", m[i])
	}
}
