# Web Framework Benchmark

## Installation

- `go get github.com/eudore/web-framework-benchmark`

## Running Benchmark

```sh
go test -bench=. github.com/eudore/web-framework-benchmark
```

环境变量NUM是加载的空中间件数量，beego没有增加加载中间件。

## Resule

时间: 2019年6月11日23点58分

基于echo使用的[测试方法][1]fork ，

echo中间件有明显性能缺陷.

iris无法运行，会panic，测试代码已注释。

```
[root@izj6cffbpd9lzl3tcm2csxz eudore]# go version
go version go1.10.1 linux/amd64
[root@izj6cffbpd9lzl3tcm2csxz web-framework-benchmark]#  go test -bench=. github.com/eudore/web-framework-benchmark
goos: linux
goarch: amd64
pkg: github.com/eudore/web-framework-benchmark
BenchmarkGolfStatic-2             	   30000	     40334 ns/op	    2432 B/op	     157 allocs/op
BenchmarkGolfGitHubAPI-2          	   30000	     56914 ns/op	    2802 B/op	     203 allocs/op
BenchmarkGolfGplusAPI-2           	  500000	      2567 ns/op	     173 B/op	      13 allocs/op
BenchmarkGolfParseAPI-2           	  300000	      4448 ns/op	     323 B/op	      26 allocs/op
BenchmarkEchoStatic-2             	   30000	     53426 ns/op	    2413 B/op	     157 allocs/op
BenchmarkEchoGitHubAPI-2          	   20000	     78275 ns/op	    2496 B/op	     203 allocs/op
BenchmarkEchoGplusAPI-2           	  300000	      4318 ns/op	     161 B/op	      13 allocs/op
BenchmarkEchoParseAPI-2           	  200000	      7493 ns/op	     381 B/op	      26 allocs/op
BenchmarkGinStatic-2              	   20000	     72672 ns/op	    8405 B/op	     157 allocs/op
BenchmarkGinGitHubAPI-2           	   10000	    110778 ns/op	   10620 B/op	     203 allocs/op
BenchmarkGinGplusAPI-2            	  300000	      5942 ns/op	     681 B/op	      13 allocs/op
BenchmarkGinParseAPI-2            	  200000	     11195 ns/op	    1421 B/op	      26 allocs/op
BenchmarkDotwebStatic-2           	    5000	    355636 ns/op	  101158 B/op	    2847 allocs/op
BenchmarkDotwebGitHubAPI-2        	    3000	    519782 ns/op	  145000 B/op	    3850 allocs/op
BenchmarkDotwebGplusAPI-2         	   50000	     30968 ns/op	    9041 B/op	     246 allocs/op
BenchmarkDotwebParseAPI-2         	   20000	     61108 ns/op	   17335 B/op	     487 allocs/op
BenchmarkBeegoStatic-2            	    5000	    285655 ns/op	   77514 B/op	     942 allocs/op
BenchmarkBeegoGitHubAPI-2         	    3000	    369264 ns/op	  104918 B/op	    1222 allocs/op
BenchmarkBeegoGplusAPI-2          	  100000	     21204 ns/op	    6437 B/op	      78 allocs/op
BenchmarkBeegoParseAPI-2          	   30000	     39594 ns/op	   12859 B/op	     156 allocs/op
BenchmarkTwigStatic-2             	   30000	     56868 ns/op	    3668 B/op	     157 allocs/op
BenchmarkTwigGitHubAPI-2          	   20000	     84616 ns/op	    4118 B/op	     203 allocs/op
BenchmarkTwigGplusAPI-2           	  300000	      4481 ns/op	     265 B/op	      13 allocs/op
BenchmarkTwigParseAPI-2           	  200000	      8303 ns/op	     589 B/op	      26 allocs/op
BenchmarkEudoreRadixStatic-2      	   20000	     78383 ns/op	     872 B/op	       0 allocs/op
BenchmarkEudoreRadixGitHubAPI-2   	   10000	    115842 ns/op	     895 B/op	       0 allocs/op
BenchmarkEudoreRadixGplusAPI-2    	  300000	      5679 ns/op	      57 B/op	       0 allocs/op
BenchmarkEudoreRadixParseAPI-2    	  200000	     10465 ns/op	     173 B/op	       0 allocs/op
BenchmarkEudoreFullStatic-2       	   20000	     78899 ns/op	     872 B/op	       0 allocs/op
BenchmarkEudoreFullGitHubAPI-2    	   10000	    115293 ns/op	     899 B/op	       0 allocs/op
BenchmarkEudoreFullGplusAPI-2     	  300000	      5602 ns/op	      57 B/op	       0 allocs/op
BenchmarkEudoreFullParseAPI-2     	  200000	     10581 ns/op	     173 B/op	       0 allocs/op
PASS
ok  	github.com/eudore/web-framework-benchmark	58.335s
[root@izj6cffbpd9lzl3tcm2csxz web-framework-benchmark]# NUM=20 go test -bench=. github.com/eudore/web-framework-benchmark
goos: linux
goarch: amd64
pkg: github.com/eudore/web-framework-benchmark
BenchmarkGolfStatic-2             	   20000	     69732 ns/op	    2156 B/op	     158 allocs/op
BenchmarkGolfGitHubAPI-2          	   20000	     99983 ns/op	    2526 B/op	     203 allocs/op
BenchmarkGolfGplusAPI-2           	  300000	      5016 ns/op	     161 B/op	      13 allocs/op
BenchmarkGolfParseAPI-2           	  200000	      9640 ns/op	     381 B/op	      26 allocs/op
BenchmarkEchoStatic-2             	   10000	    209040 ns/op	   52376 B/op	    3297 allocs/op
BenchmarkEchoGitHubAPI-2          	    5000	    293984 ns/op	   67484 B/op	    4263 allocs/op
BenchmarkEchoGplusAPI-2           	  100000	     17197 ns/op	    4351 B/op	     273 allocs/op
BenchmarkEchoParseAPI-2           	   50000	     34017 ns/op	    8702 B/op	     546 allocs/op
BenchmarkGinStatic-2              	   20000	     89798 ns/op	    8406 B/op	     157 allocs/op
BenchmarkGinGitHubAPI-2           	   10000	    131699 ns/op	   10624 B/op	     203 allocs/op
BenchmarkGinGplusAPI-2            	  200000	      7582 ns/op	     710 B/op	      13 allocs/op
BenchmarkGinParseAPI-2            	  100000	     14377 ns/op	    1421 B/op	      26 allocs/op
BenchmarkDotwebStatic-2           	    5000	    372720 ns/op	  101271 B/op	    2847 allocs/op
BenchmarkDotwebGitHubAPI-2        	    3000	    598605 ns/op	  144713 B/op	    3847 allocs/op
BenchmarkDotwebGplusAPI-2         	   50000	     33761 ns/op	    8998 B/op	     246 allocs/op
BenchmarkDotwebParseAPI-2         	   20000	     62377 ns/op	   17365 B/op	     487 allocs/op
BenchmarkBeegoStatic-2            	    5000	    272420 ns/op	   77514 B/op	     942 allocs/op
BenchmarkBeegoGitHubAPI-2         	    3000	    357711 ns/op	  104917 B/op	    1222 allocs/op
BenchmarkBeegoGplusAPI-2          	  100000	     21140 ns/op	    6437 B/op	      78 allocs/op
BenchmarkBeegoParseAPI-2          	   30000	     41504 ns/op	   12859 B/op	     156 allocs/op
BenchmarkTwigStatic-2             	   30000	     58564 ns/op	    3668 B/op	     157 allocs/op
BenchmarkTwigGitHubAPI-2          	   20000	     84043 ns/op	    4118 B/op	     203 allocs/op
BenchmarkTwigGplusAPI-2           	  300000	      4594 ns/op	     265 B/op	      13 allocs/op
BenchmarkTwigParseAPI-2           	  200000	      8460 ns/op	     589 B/op	      26 allocs/op
BenchmarkEudoreRadixStatic-2      	   20000	     92960 ns/op	     873 B/op	       0 allocs/op
BenchmarkEudoreRadixGitHubAPI-2   	   10000	    135900 ns/op	     900 B/op	       0 allocs/op
BenchmarkEudoreRadixGplusAPI-2    	  200000	      6920 ns/op	      86 B/op	       0 allocs/op
BenchmarkEudoreRadixParseAPI-2    	  100000	     13049 ns/op	     173 B/op	       0 allocs/op
BenchmarkEudoreFullStatic-2       	   20000	     92802 ns/op	     874 B/op	       0 allocs/op
BenchmarkEudoreFullGitHubAPI-2    	   10000	    140953 ns/op	     904 B/op	       0 allocs/op
BenchmarkEudoreFullGplusAPI-2     	  200000	      6850 ns/op	      86 B/op	       0 allocs/op
BenchmarkEudoreFullParseAPI-2     	  100000	     12480 ns/op	     173 B/op	       0 allocs/op
PASS
ok  	github.com/eudore/web-framework-benchmark	60.150s
```

[1]: https://github.com/vishr/web-framework-benchmark
