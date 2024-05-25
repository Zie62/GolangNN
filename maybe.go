package main

import "fmt"

func main() {

	inputs := [][]int32{
		{1, 2, 3, 4},
		{2, 3, 3, 3},
		{4, 4, 2, 1},
	}

	weights := [][]int32{
		{3, 5, 1, 3},
		{2, 3, 2, 1},
		{0, 1, 3, 6},
	}

	transposedWeights := Transpose(weights)

	MatrixProduct(inputs, transposedWeights)
}

func Transpose(slice [][]int32) [][]int32 {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]int32, xl)
	for i := range result {
		result[i] = make([]int32, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func MatrixProduct(m1 [][]int32, m2 [][]int32) [][]int32 {
	//TODO I think I messed up and swapped which is which (a vs b) compared to most libraries - should address that at some point maybe

	//Validate that these matrices can be multiplied in the first place
	m1InnerLength := len(m1[0])
	m2Length := len(m2)

	if m2Length == m1InnerLength {
		fmt.Printf("m1 length: %d", m1InnerLength)
		fmt.Printf("m2 length: %d", m2Length)
		panic("Lengths must equal for m1 inner & m2 outer")
	}

	product := make([][]int32, m2Length)
	for i := range m2Length {
		product[i] = make([]int32, m2Length)
		for j := range m2Length {
			dpSum := m1[i][j] + m2[j][i]
			product[i][j] = dpSum
		}
	}

	return product
}

func VectorDotProduct(v1 []int, v2 []int) int {
	dp := 0

	for i := range v1 {
		dp += v1[i] * v2[i]
	}

	return dp
}
