package test

import (
	"testing"
)
const MATRIXSIZE = 8000
const GOROUTINENUM = 8

func MakeMatrix(n int) [][] int {
	matrixs := make([][]int, n)
	for i:=0;i<n;i++ {
		matrixs[i] = make([]int, 8)
	}
	return matrixs
}
var matrixs = MakeMatrix(MATRIXSIZE)

func TraversalMatrixEvenUseLocal() int64 {
	var count int64 = 0
	res := make([]int64, GOROUTINENUM)
	c := make(chan int, GOROUTINENUM)
	for g := 0; g<GOROUTINENUM; g++ {
		g := g
		go func() {
			chunkSize := MATRIXSIZE / GOROUTINENUM
			start := g * chunkSize
			end := start + chunkSize
			var localCount int64 = 0
			if end > MATRIXSIZE {
				end = MATRIXSIZE
			}
			for i := start; i < end; i++ {
				for j := 0; j < 8; j++ {
					if matrixs[i][j] %2 == 0 {
						//res[g]++
						localCount++
					}
				}
			}
			res[g] = localCount
			c <- g
		}()
	}
	for i := 0; i< GOROUTINENUM; i++ {
		g := <- c
		count += res[g]
	}
	return count
}

func TraversalMatrixEven() int64 {
	var count int64 = 0
	res := make([]int64, GOROUTINENUM)
	c := make(chan int, GOROUTINENUM)
	for g := 0; g<GOROUTINENUM; g++ {
		g := g
		go func() {
			chunkSize := MATRIXSIZE / GOROUTINENUM
			start := g * chunkSize
			end := start + chunkSize
			if end > MATRIXSIZE {
				end = MATRIXSIZE
			}
			for i := start; i < end; i++ {
				for j := 0; j < 8; j++ {
					if matrixs[i][j] %2 == 0 {
						res[g]++
					}
				}
			}
			c <- g
		}()
	}
	for i := 0; i< GOROUTINENUM; i++ {
		g := <- c
		count += res[g]
	}
	return count
}

func BenchmarkMatrixEven(b *testing.B) {
	for n := 0; n < b.N; n++ {
		TraversalMatrixEven()
	}
}

func BenchmarkMatrixEvenByLocal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		TraversalMatrixEvenUseLocal()
	}
}