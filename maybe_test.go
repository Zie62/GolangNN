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
