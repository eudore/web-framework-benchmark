# Web Framework Benchmark

## Installation

- `go get github.com/eudore/web-framework-benchmark`

## Running Benchmark

```sh
go test -bench=. github.com/eudore/web-framework-benchmark
```

环境变量NUM是加载的空中间件数量，beego没有增加加载中间件。

## Resule

时间: 2019年3月17日17:34

基于echo使用的[测试方法][1]fork ，

echo中间件有明显性能缺陷.

iris无法运行，会panic，测试代码已注释。

```
[root@izj6cffbpd9lzl3tcm2csxz eudore]# go version
go version go1.10.1 linux/amd64
[root@izj6cffbpd9lzl3tcm2csxz web-framework-benchmark]# go test -bench=. github.com/eudore/web-framework-benchmark
goos: linux
goarch: amd64
pkg: github.com/eudore/web-framework-benchmark
BenchmarkGolfStatic-2             	   30000	     40388 ns/op	    2432 B/op	     157 allocs/op
BenchmarkGolfGitHubAPI-2          	   20000	     57603 ns/op	    2526 B/op	     203 allocs/op
BenchmarkGolfGplusAPI-2           	  500000	      2576 ns/op	     173 B/op	      13 allocs/op
BenchmarkGolfParseAPI-2           	  300000	      4438 ns/op	     323 B/op	      26 allocs/op
BenchmarkEchoStatic-2             	   30000	     64028 ns/op	    2413 B/op	     157 allocs/op
BenchmarkEchoGitHubAPI-2          	   20000	     75715 ns/op	    2496 B/op	     203 allocs/op
BenchmarkEchoGplusAPI-2           	  300000	      4131 ns/op	     161 B/op	      13 allocs/op
BenchmarkEchoParseAPI-2           	  200000	      7222 ns/op	     381 B/op	      26 allocs/op
BenchmarkGinStatic-2              	   20000	     67784 ns/op	    8405 B/op	     157 allocs/op
BenchmarkGinGitHubAPI-2           	   20000	     92765 ns/op	   10614 B/op	     203 allocs/op
BenchmarkGinGplusAPI-2            	  200000	      5434 ns/op	     710 B/op	      13 allocs/op
BenchmarkGinParseAPI-2            	  200000	     10485 ns/op	    1421 B/op	      26 allocs/op
BenchmarkDotwebStatic-2           	    3000	    485627 ns/op	  101132 B/op	    2849 allocs/op
BenchmarkDotwebGitHubAPI-2        	    2000	    619509 ns/op	  144710 B/op	    3852 allocs/op
BenchmarkDotwebGplusAPI-2         	   50000	     33663 ns/op	    9025 B/op	     246 allocs/op
BenchmarkDotwebParseAPI-2         	   20000	     74080 ns/op	   17333 B/op	     487 allocs/op
BenchmarkBeegoStatic-2            	    2000	    766096 ns/op	  145630 B/op	    2984 allocs/op
BenchmarkBeegoGitHubAPI-2         	    2000	   1001140 ns/op	  194946 B/op	    3864 allocs/op
BenchmarkBeegoGplusAPI-2          	   20000	     68075 ns/op	   12066 B/op	     247 allocs/op
BenchmarkBeegoParseAPI-2          	   10000	    122213 ns/op	   24177 B/op	     494 allocs/op
BenchmarkTwigStatic-2             	   30000	     58957 ns/op	    3668 B/op	     157 allocs/op
BenchmarkTwigGitHubAPI-2          	   20000	     83141 ns/op	    4118 B/op	     203 allocs/op
BenchmarkTwigGplusAPI-2           	  300000	      4566 ns/op	     265 B/op	      13 allocs/op
BenchmarkTwigParseAPI-2           	  200000	      8476 ns/op	     589 B/op	      26 allocs/op
BenchmarkEudoreRadixStatic-2      	   20000	     68244 ns/op	    1030 B/op	       0 allocs/op
BenchmarkEudoreRadixGitHubAPI-2   	   20000	     92013 ns/op	    1831 B/op	       0 allocs/op
BenchmarkEudoreRadixGplusAPI-2    	  300000	      4645 ns/op	     258 B/op	       0 allocs/op
BenchmarkEudoreRadixParseAPI-2    	  200000	      8884 ns/op	     401 B/op	       0 allocs/op
BenchmarkEudoreFullStatic-2       	   20000	     75082 ns/op	    1030 B/op	       0 allocs/op
BenchmarkEudoreFullGitHubAPI-2    	   10000	    102973 ns/op	    1850 B/op	       0 allocs/op
BenchmarkEudoreFullGplusAPI-2     	  300000	      4972 ns/op	     258 B/op	       0 allocs/op
BenchmarkEudoreFullParseAPI-2     	  200000	     12942 ns/op	     401 B/op	       0 allocs/op
PASS
ok  	github.com/eudore/web-framework-benchmark	60.653s
[root@izj6cffbpd9lzl3tcm2csxz web-framework-benchmark]# NUM=20 go test -bench=. github.com/eudore/web-framework-benchmark
goos: linux
goarch: amd64
pkg: github.com/eudore/web-framework-benchmark
BenchmarkGolfStatic-2             	   20000	     67996 ns/op	    2156 B/op	     158 allocs/op
BenchmarkGolfGitHubAPI-2          	   20000	     96828 ns/op	    2526 B/op	     203 allocs/op
BenchmarkGolfGplusAPI-2           	  300000	      4796 ns/op	     161 B/op	      13 allocs/op
BenchmarkGolfParseAPI-2           	  200000	      8715 ns/op	     381 B/op	      26 allocs/op
BenchmarkEchoStatic-2             	   10000	    212161 ns/op	   52376 B/op	    3297 allocs/op
BenchmarkEchoGitHubAPI-2          	    5000	    317832 ns/op	   67482 B/op	    4263 allocs/op
BenchmarkEchoGplusAPI-2           	  100000	     17860 ns/op	    4351 B/op	     273 allocs/op
BenchmarkEchoParseAPI-2           	   50000	     35064 ns/op	    8702 B/op	     546 allocs/op
BenchmarkGinStatic-2              	   20000	     85640 ns/op	    8406 B/op	     157 allocs/op
BenchmarkGinGitHubAPI-2           	   10000	    112078 ns/op	   10623 B/op	     203 allocs/op
BenchmarkGinGplusAPI-2            	  200000	      6603 ns/op	     710 B/op	      13 allocs/op
BenchmarkGinParseAPI-2            	  100000	     12891 ns/op	    1421 B/op	      26 allocs/op
BenchmarkDotwebStatic-2           	    3000	    447462 ns/op	  100768 B/op	    2844 allocs/op
BenchmarkDotwebGitHubAPI-2        	    2000	    634951 ns/op	  144210 B/op	    3847 allocs/op
BenchmarkDotwebGplusAPI-2         	   50000	     32621 ns/op	    9011 B/op	     246 allocs/op
BenchmarkDotwebParseAPI-2         	   20000	     68279 ns/op	   17363 B/op	     487 allocs/op
BenchmarkBeegoStatic-2            	    2000	    808517 ns/op	  145629 B/op	    2984 allocs/op
BenchmarkBeegoGitHubAPI-2         	    2000	   1039742 ns/op	  194946 B/op	    3864 allocs/op
BenchmarkBeegoGplusAPI-2          	   20000	     60741 ns/op	   12066 B/op	     247 allocs/op
BenchmarkBeegoParseAPI-2          	   10000	    126334 ns/op	   24177 B/op	     494 allocs/op
BenchmarkTwigStatic-2             	   30000	     60154 ns/op	    3668 B/op	     157 allocs/op
BenchmarkTwigGitHubAPI-2          	   20000	     85172 ns/op	    4118 B/op	     203 allocs/op
BenchmarkTwigGplusAPI-2           	  300000	      4678 ns/op	     265 B/op	      13 allocs/op
BenchmarkTwigParseAPI-2           	  200000	      8550 ns/op	     589 B/op	      26 allocs/op
BenchmarkEudoreRadixStatic-2      	   10000	    100496 ns/op	    1039 B/op	       0 allocs/op
BenchmarkEudoreRadixGitHubAPI-2   	   10000	    119274 ns/op	    1849 B/op	       0 allocs/op
BenchmarkEudoreRadixGplusAPI-2    	  300000	      5638 ns/op	     258 B/op	       0 allocs/op
BenchmarkEudoreRadixParseAPI-2    	  200000	     11082 ns/op	     401 B/op	       0 allocs/op
BenchmarkEudoreFullStatic-2       	   20000	     92407 ns/op	    1032 B/op	       0 allocs/op
BenchmarkEudoreFullGitHubAPI-2    	   10000	    124795 ns/op	    1854 B/op	       0 allocs/op
BenchmarkEudoreFullGplusAPI-2     	  200000	      6136 ns/op	     194 B/op	       0 allocs/op
BenchmarkEudoreFullParseAPI-2     	  200000	     11834 ns/op	     401 B/op	       0 allocs/op
PASS
ok  	github.com/eudore/web-framework-benchmark	59.173s
```

[1]: https://github.com/vishr/web-framework-benchmark
