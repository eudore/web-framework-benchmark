# Web Framework Benchmark

## Installation

- `go get github.com/eudore/web-framework-benchmark`

## Running Benchmark

```sh
go test -bench=.
```

## Resule

时间：2019年2月23日22:37:24

基于echo[测试方法][1]fork ，

iris无法运行

```
goos: linux
goarch: amd64
pkg: github.com/eudore/web-framework-benchmark
BenchmarkGolfStatic-2        	   30000	     40431 ns/op	    2432 B/op	     157 allocs/op
BenchmarkGolfGitHubAPI-2     	   30000	     60360 ns/op	    2802 B/op	     203 allocs/op
BenchmarkGolfGplusAPI-2      	  500000	      2687 ns/op	     173 B/op	      13 allocs/op
BenchmarkGolfParseAPI-2      	  300000	      5274 ns/op	     323 B/op	      26 allocs/op
BenchmarkEchoStatic-2        	   30000	     51219 ns/op	    2413 B/op	     157 allocs/op
BenchmarkEchoGitHubAPI-2     	   20000	     77637 ns/op	    2496 B/op	     203 allocs/op
BenchmarkEchoGplusAPI-2      	  300000	      4121 ns/op	     161 B/op	      13 allocs/op
BenchmarkEchoParseAPI-2      	  200000	      7501 ns/op	     381 B/op	      26 allocs/op
BenchmarkGinStatic-2         	   20000	     66083 ns/op	    8404 B/op	     157 allocs/op
BenchmarkGinGitHubAPI-2      	   20000	     91717 ns/op	   10614 B/op	     203 allocs/op
BenchmarkGinGplusAPI-2       	  300000	      5484 ns/op	     681 B/op	      13 allocs/op
BenchmarkGinParseAPI-2       	  200000	     10125 ns/op	    1421 B/op	      26 allocs/op
BenchmarkDotwebStatic-2      	    5000	    369653 ns/op	  101315 B/op	    2847 allocs/op
BenchmarkDotwebGitHubAPI-2   	    3000	    521054 ns/op	  144996 B/op	    3849 allocs/op
BenchmarkDotwebGplusAPI-2    	   50000	     34357 ns/op	    9041 B/op	     246 allocs/op
BenchmarkDotwebParseAPI-2    	   20000	     66955 ns/op	   17339 B/op	     487 allocs/op
BenchmarkBeegoStatic-2       	    2000	    743984 ns/op	  145627 B/op	    2984 allocs/op
BenchmarkBeegoGitHubAPI-2    	    2000	   1202679 ns/op	  194944 B/op	    3864 allocs/op
BenchmarkBeegoGplusAPI-2     	   20000	     72291 ns/op	   12066 B/op	     247 allocs/op
BenchmarkBeegoParseAPI-2     	   10000	    120126 ns/op	   24176 B/op	     494 allocs/op
BenchmarkTwigStatic-2        	   30000	     55410 ns/op	    2413 B/op	     157 allocs/op
BenchmarkTwigGitHubAPI-2     	   20000	     80728 ns/op	    2495 B/op	     203 allocs/op
BenchmarkTwigGplusAPI-2      	  300000	      4494 ns/op	     161 B/op	      13 allocs/op
BenchmarkTwigParseAPI-2      	  200000	      7837 ns/op	     381 B/op	      26 allocs/op
BenchmarkEudoreStatic-2      	   30000	     59743 ns/op	    1156 B/op	       0 allocs/op
BenchmarkEudoreGitHubAPI-2   	   20000	     61692 ns/op	       9 B/op	       0 allocs/op
BenchmarkEudoreGplusAPI-2    	  500000	      2826 ns/op	       0 B/op	       0 allocs/op
BenchmarkEudoreParseAPI-2    	  300000	      5013 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/eudore/web-framework-benchmark	53.981s
```

[1]: https://github.com/eudore/web-framework-benchmark
