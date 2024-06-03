package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectorDotProduct(t *testing.T) {
	v1 := []int{1, 2, 3, 4}
	v2 := []int{2, 3, 1, 4}

	dp := VectorDotProduct(v1, v2)

	assert.Equal(t, 27, dp, "Dot product did not produce correct output")
}

func TestTranspose(t *testing.T) {
	weights := [][]int{
		{3, 5, 1, 3},
		{2, 3, 2, 1},
		{0, 1, 3, 6},
	}

	//TODO: Validate the transpose function works as intended
	transposedWeights := Transpose(weights)

	if transposedWeights == nil {
		panic("ahhhhhhh")
	}
}

func TestMatrixProduct(t *testing.T) {
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

	//TODO: Validate the product we get out contains the expected values & shape
	product := MatrixProduct(inputs, transposedWeights)

	//We should get a 3x3 since we are multiplying a 3x4 and 4x3 matrix
	assert.Equal(t, 3, len(product))
	assert.Equal(t, 3, len(product[0]))

	//Validate a couple of specific values for this test

	if product == nil {
		panic("ahhhhhhh")
	}
}
