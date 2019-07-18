# Web Framework Benchmark

## Installation

- `go get github.com/eudore/web-framework-benchmark`

## Running Benchmark

```sh
go test -bench=. github.com/eudore/web-framework-benchmark
```

环境变量NUM是加载的空中间件数量，beego等没有增加加载中间件。

## Resule

时间: 2019年7月18日21点28分

基于echo使用的[测试方法][1]fork ，

iris无法运行，会panic，测试代码已注释。

```
[root@izj6cffbpd9lzl3tcm2csxz eudore]# go version
go version go1.10.1 linux/amd64
[root@izj6cffbpd9lzl3tcm2csxz web-framework-benchmark]# go test -bench=. github.com/eudore/web-framework-benchmark
goos: linux
goarch: amd64
pkg: github.com/eudore/web-framework-benchmark
BenchmarkGolfStatic-2              	   50000	     39422 ns/op	    1961 B/op	     157 allocs/op
BenchmarkGolfGitHubAPI-2           	   30000	     56541 ns/op	    2802 B/op	     203 allocs/op
BenchmarkGolfGplusAPI-2            	  500000	      2552 ns/op	     173 B/op	      13 allocs/op
BenchmarkGolfParseAPI-2            	  300000	      4435 ns/op	     323 B/op	      26 allocs/op
BenchmarkEchoStatic-2              	   30000	     53307 ns/op	    2413 B/op	     157 allocs/op
BenchmarkEchoGitHubAPI-2           	   20000	     77366 ns/op	    2496 B/op	     203 allocs/op
BenchmarkEchoGplusAPI-2            	  300000	      4235 ns/op	     161 B/op	      13 allocs/op
BenchmarkEchoParseAPI-2            	  200000	      7160 ns/op	     381 B/op	      26 allocs/op
BenchmarkGinStatic-2               	   20000	     76720 ns/op	    8405 B/op	     157 allocs/op
BenchmarkGinGitHubAPI-2            	   10000	    112871 ns/op	   10620 B/op	     203 allocs/op
BenchmarkGinGplusAPI-2             	  200000	      6371 ns/op	     710 B/op	      13 allocs/op
BenchmarkGinParseAPI-2             	  200000	     11692 ns/op	    1421 B/op	      26 allocs/op
BenchmarkDotwebStatic-2            	    5000	    396412 ns/op	  100932 B/op	    2844 allocs/op
BenchmarkDotwebGitHubAPI-2         	    3000	    543539 ns/op	  145204 B/op	    3850 allocs/op
BenchmarkDotwebGplusAPI-2          	   50000	     33834 ns/op	    9024 B/op	     246 allocs/op
BenchmarkDotwebParseAPI-2          	   20000	     69214 ns/op	   17341 B/op	     487 allocs/op
BenchmarkBeegoStatic-2             	    5000	    274688 ns/op	   77514 B/op	     942 allocs/op
BenchmarkBeegoGitHubAPI-2          	    3000	    381696 ns/op	  104917 B/op	    1222 allocs/op
BenchmarkBeegoGplusAPI-2           	  100000	     23781 ns/op	    6437 B/op	      78 allocs/op
BenchmarkBeegoParseAPI-2           	   30000	     40864 ns/op	   12859 B/op	     156 allocs/op
BenchmarkTwigStatic-2              	   30000	     56564 ns/op	    3668 B/op	     157 allocs/op
BenchmarkTwigGitHubAPI-2           	   10000	    111532 ns/op	    4122 B/op	     203 allocs/op
BenchmarkTwigGplusAPI-2            	  300000	      4543 ns/op	     265 B/op	      13 allocs/op
BenchmarkTwigParseAPI-2            	  200000	      8316 ns/op	     589 B/op	      26 allocs/op
BenchmarkEudoreRadixStatic-2       	   20000	     77580 ns/op	     872 B/op	       0 allocs/op
BenchmarkEudoreRadixGitHubAPI-2    	   10000	    119841 ns/op	     895 B/op	       0 allocs/op
BenchmarkEudoreRadixGplusAPI-2     	  300000	      5837 ns/op	      57 B/op	       0 allocs/op
BenchmarkEudoreRadixParseAPI-2     	  200000	     10479 ns/op	     173 B/op	       0 allocs/op
BenchmarkEudoreFullStatic-2        	   20000	     76923 ns/op	     872 B/op	       0 allocs/op
BenchmarkEudoreFullGitHubAPI-2     	   10000	    112567 ns/op	     899 B/op	       0 allocs/op
BenchmarkEudoreFullGplusAPI-2      	  300000	      5623 ns/op	      57 B/op	       0 allocs/op
BenchmarkEudoreFullParseAPI-2      	  200000	     10473 ns/op	     173 B/op	       0 allocs/op
BenchmarkHttprouterStatic-2        	   50000	     24124 ns/op	    1949 B/op	     157 allocs/op
BenchmarkHttprouterGitHubAPI-2     	   30000	     55458 ns/op	   16571 B/op	     370 allocs/op
BenchmarkHttprouterGplusAPI-2      	  500000	      2552 ns/op	     813 B/op	      24 allocs/op
BenchmarkHttprouterParseAPI-2      	  500000	      3791 ns/op	     986 B/op	      42 allocs/op
BenchmarkErouterRadixStatic-2      	   30000	     43719 ns/op	    2412 B/op	     157 allocs/op
BenchmarkErouterRadixGitHubAPI-2   	   20000	     61982 ns/op	    2501 B/op	     203 allocs/op
BenchmarkErouterRadixGplusAPI-2    	  500000	      2534 ns/op	     173 B/op	      13 allocs/op
BenchmarkErouterRadixParseAPI-2    	  300000	      4590 ns/op	     323 B/op	      26 allocs/op
BenchmarkErouterFullStatic-2       	   30000	     44715 ns/op	    2413 B/op	     157 allocs/op
BenchmarkErouterFullGitHubAPI-2    	   20000	     63466 ns/op	    2503 B/op	     203 allocs/op
BenchmarkErouterFullGplusAPI-2     	  500000	      2575 ns/op	     173 B/op	      13 allocs/op
BenchmarkErouterFullParseAPI-2     	  300000	      4608 ns/op	     323 B/op	      26 allocs/op
PASS
ok  	github.com/eudore/web-framework-benchmark	78.223s
[root@izj6cffbpd9lzl3tcm2csxz web-framework-benchmark]# NUM=20 go test -bench=. github.com/eudore/web-framework-benchmark
goos: linux
goarch: amd64
pkg: github.com/eudore/web-framework-benchmark
BenchmarkGolfStatic-2              	   20000	     69660 ns/op	    2156 B/op	     158 allocs/op
BenchmarkGolfGitHubAPI-2           	   20000	     99094 ns/op	    2526 B/op	     203 allocs/op
BenchmarkGolfGplusAPI-2            	  300000	      4998 ns/op	     161 B/op	      13 allocs/op
BenchmarkGolfParseAPI-2            	  200000	      9618 ns/op	     381 B/op	      26 allocs/op
BenchmarkEchoStatic-2              	   10000	    234609 ns/op	   52376 B/op	    3297 allocs/op
BenchmarkEchoGitHubAPI-2           	    5000	    290072 ns/op	   67483 B/op	    4263 allocs/op
BenchmarkEchoGplusAPI-2            	  100000	     17354 ns/op	    4351 B/op	     273 allocs/op
BenchmarkEchoParseAPI-2            	   50000	     33383 ns/op	    8702 B/op	     546 allocs/op
BenchmarkGinStatic-2               	   20000	    108324 ns/op	    8406 B/op	     157 allocs/op
BenchmarkGinGitHubAPI-2            	   10000	    135712 ns/op	   10624 B/op	     203 allocs/op
BenchmarkGinGplusAPI-2             	  200000	      7522 ns/op	     710 B/op	      13 allocs/op
BenchmarkGinParseAPI-2             	  100000	     15113 ns/op	    1421 B/op	      26 allocs/op
BenchmarkDotwebStatic-2            	    3000	    415992 ns/op	  100845 B/op	    2846 allocs/op
BenchmarkDotwebGitHubAPI-2         	    2000	    680612 ns/op	  144564 B/op	    3848 allocs/op
BenchmarkDotwebGplusAPI-2          	   50000	     35301 ns/op	    9013 B/op	     246 allocs/op
BenchmarkDotwebParseAPI-2          	   20000	     66566 ns/op	   17363 B/op	     487 allocs/op
BenchmarkBeegoStatic-2             	    5000	    273241 ns/op	   77514 B/op	     942 allocs/op
BenchmarkBeegoGitHubAPI-2          	    3000	    379682 ns/op	  104917 B/op	    1222 allocs/op
BenchmarkBeegoGplusAPI-2           	  100000	     21370 ns/op	    6437 B/op	      78 allocs/op
BenchmarkBeegoParseAPI-2           	   30000	     40965 ns/op	   12859 B/op	     156 allocs/op
BenchmarkTwigStatic-2              	   30000	     56689 ns/op	    3668 B/op	     157 allocs/op
BenchmarkTwigGitHubAPI-2           	   20000	     82984 ns/op	    4118 B/op	     203 allocs/op
BenchmarkTwigGplusAPI-2            	  300000	      4503 ns/op	     265 B/op	      13 allocs/op
BenchmarkTwigParseAPI-2            	  200000	      8303 ns/op	     589 B/op	      26 allocs/op
BenchmarkEudoreRadixStatic-2       	   20000	     92419 ns/op	     873 B/op	       0 allocs/op
BenchmarkEudoreRadixGitHubAPI-2    	   10000	    133850 ns/op	     900 B/op	       0 allocs/op
BenchmarkEudoreRadixGplusAPI-2     	  200000	      6924 ns/op	      86 B/op	       0 allocs/op
BenchmarkEudoreRadixParseAPI-2     	  100000	     13067 ns/op	     173 B/op	       0 allocs/op
BenchmarkEudoreFullStatic-2        	   20000	     94141 ns/op	     874 B/op	       0 allocs/op
BenchmarkEudoreFullGitHubAPI-2     	   10000	    133557 ns/op	     904 B/op	       0 allocs/op
BenchmarkEudoreFullGplusAPI-2      	  200000	      7022 ns/op	      86 B/op	       0 allocs/op
BenchmarkEudoreFullParseAPI-2      	  100000	     13122 ns/op	     173 B/op	       0 allocs/op
BenchmarkHttprouterStatic-2        	   50000	     24855 ns/op	    1949 B/op	     157 allocs/op
BenchmarkHttprouterGitHubAPI-2     	   30000	     54699 ns/op	   16571 B/op	     370 allocs/op
BenchmarkHttprouterGplusAPI-2      	  500000	      2668 ns/op	     813 B/op	      24 allocs/op
BenchmarkHttprouterParseAPI-2      	  500000	      4278 ns/op	     986 B/op	      42 allocs/op
BenchmarkErouterRadixStatic-2      	   20000	     75091 ns/op	    2129 B/op	     157 allocs/op
BenchmarkErouterRadixGitHubAPI-2   	   10000	    102855 ns/op	    2521 B/op	     203 allocs/op
BenchmarkErouterRadixGplusAPI-2    	  300000	      4866 ns/op	     161 B/op	      13 allocs/op
BenchmarkErouterRadixParseAPI-2    	  200000	      9675 ns/op	     381 B/op	      26 allocs/op
BenchmarkErouterFullStatic-2       	   20000	     72833 ns/op	    2129 B/op	     157 allocs/op
BenchmarkErouterFullGitHubAPI-2    	   10000	    107174 ns/op	    2525 B/op	     203 allocs/op
BenchmarkErouterFullGplusAPI-2     	  300000	      5156 ns/op	     161 B/op	      13 allocs/op
BenchmarkErouterFullParseAPI-2     	  200000	     10138 ns/op	     381 B/op	      26 allocs/op
PASS
ok  	github.com/eudore/web-framework-benchmark	81.163s
```
