package test

import (
	"testing"
)

func TraversalMatrixSumsByRow() int {
	sum := 0
	for i:= 0; i< MATRIXSIZE; i++ {
		for j := 0; j<8; j++ {
			sum += matrixs[i][j]
		}
	}
	return sum
}

func TraversalMatrixSumsByColumn() int {
	sum := 0
	for i:= 0; i< 8; i++ {
		for j := 0; j<MATRIXSIZE; j++ {
			sum += matrixs[j][i]
		}
	}
	return sum
}

func BenchmarkTraversalSumByRow(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		TraversalMatrixSumsByRow()
	}
}

func BenchmarkTraversalSumByColumn(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		TraversalMatrixSumsByColumn()
	}
}