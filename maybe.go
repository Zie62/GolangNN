package main

import "fmt"

func main() {

	inputs := [][]int{
		{1, 2, 3, 4},
		{2, 3, 3, 3},
		{4, 4, 2, 1},
	}

	weights := [][]int{
		{3, 5, 1, 3},
		{2, 3, 2, 1},
		{0, 1, 3, 6},
	}

	transposedWeights := Transpose(weights)

	MatrixProduct(inputs, transposedWeights)
}

func Transpose(slice [][]int) [][]int {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]int, xl)
	for i := range result {
		result[i] = make([]int, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func MatrixProduct(m1 [][]int, m2 [][]int) [][]int {
	//TODO I think I messed up and swapped which is which (a vs b) compared to most libraries - should address that at some point maybe

	//Validate that these matrices can be multiplied in the first place
	m1Columns := len(m1[0])
	m2Rows := len(m2)

	//Get the lengths of our output matrix
	// m1Rows := len(m1) //Output rows count
	// m2Columns := len(m2[0]) //Output columns count (number of items in each row)

	if m1Columns != m2Rows {
		fmt.Printf("m1 length: %d", m1Columns)
		fmt.Printf("m2 length: %d", m2Rows)
		panic("Lengths must equal for m1 inner (columns) & m2 outer (rows)")
	}

	//TODO: Fix this to output a properly sized matrix - currently is not doing that
	product := make([][]int, m2Rows)
	for i := range m2Rows {
		product[i] = make([]int, m2Rows)
		for j := range m2Rows {
			dpSum := m1[i][j] * m2[j][i]
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

type Node struct {
	Bias    int
	Inputs  Matrix
	Weights Matrix
}

type Matrix struct {
	Data [][]int
}
