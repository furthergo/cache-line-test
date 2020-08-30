# cache-line-test

# usage

```shell script
cd test && go test -bench=.
```

Out put is
```shell script
 BenchmarkMatrixEven-12              	   26013	     45509 ns/op
 BenchmarkMatrixEvenByLocal-12       	   46527	     25937 ns/op
 BenchmarkTraversalSumByRow-12       	   22184	     53765 ns/op
 BenchmarkTraversalSumByColumn-12    	   12649	     94891 ns/op
 PASS
 ok  	github.com/futhergo/cache-line-test/test	7.033s
```
