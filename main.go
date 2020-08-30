package main

import "testing"

const SizeNum = 1000

func GenerateMatrix() [][]int {
	matrix := make([][]int, SizeNum)
	for i, _ := range matrix {
		matrix[i] = make([]int, SizeNum)
	}
	return matrix
}

func traversalMatrix(m [][]int) int {
	sum := 0
	for i := 0; i < SizeNum; i++ {
		for j := 0; j < SizeNum; j++ {
			sum += m[i][j]
		}
	}
	return sum
}

func BenchmarkTraversalMatrix(test *testing.B) {
	for i:=0 ;i<test.N; i++ {
		m := GenerateMatrix()
		traversalMatrix(m)
	}
}

func main() {

}
