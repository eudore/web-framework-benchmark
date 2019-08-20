# Web Framework Benchmark

## Installation

- `go get github.com/eudore/web-framework-benchmark`

## Running Benchmark

```sh
go test -bench=. github.com/eudore/web-framework-benchmark
```

环境变量NUM是加载的空中间件数量，数据为空就是没有增加加载中间件。

## Resule

时间: 2019年8月20日

基于echo使用的[测试方法](https://github.com/vishr/web-framework-benchmark)fork ，

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
BenchmarkGolfStatic-2              	   20000	     71483 ns/op	    2156 B/op	     158 allocs/op
BenchmarkGolfGitHubAPI-2           	   20000	    101679 ns/op	    2526 B/op	     203 allocs/op
BenchmarkGolfGplusAPI-2            	  300000	      5082 ns/op	     161 B/op	      13 allocs/op
BenchmarkGolfParseAPI-2            	  200000	      9094 ns/op	     381 B/op	      26 allocs/op
BenchmarkEchoStatic-2              	   10000	    213412 ns/op	   52376 B/op	    3297 allocs/op
BenchmarkEchoGitHubAPI-2           	    5000	    320824 ns/op	   67483 B/op	    4263 allocs/op
BenchmarkEchoGplusAPI-2            	  100000	     17658 ns/op	    4351 B/op	     273 allocs/op
BenchmarkEchoParseAPI-2            	   50000	     34476 ns/op	    8702 B/op	     546 allocs/op
BenchmarkGinStatic-2               	   20000	     99110 ns/op	    8406 B/op	     157 allocs/op
BenchmarkGinGitHubAPI-2            	   10000	    143052 ns/op	   10624 B/op	     203 allocs/op
BenchmarkGinGplusAPI-2             	  200000	      7677 ns/op	     710 B/op	      13 allocs/op
BenchmarkGinParseAPI-2             	  100000	     14407 ns/op	    1421 B/op	      26 allocs/op
BenchmarkDotwebStatic-2            	    3000	    403247 ns/op	  101168 B/op	    2848 allocs/op
BenchmarkDotwebGitHubAPI-2         	    2000	    596129 ns/op	  144462 B/op	    3847 allocs/op
BenchmarkDotwebGplusAPI-2          	   50000	     32538 ns/op	    9019 B/op	     246 allocs/op
BenchmarkDotwebParseAPI-2          	   20000	     71272 ns/op	   17397 B/op	     487 allocs/op
BenchmarkBeegoStatic-2             	2000000000	         0.00 ns/op
BenchmarkBeegoGitHubAPI-2          	2000000000	         0.00 ns/op
BenchmarkBeegoGplusAPI-2           	2000000000	         0.00 ns/op
BenchmarkBeegoParseAPI-2           	2000000000	         0.00 ns/op
BenchmarkTwigStatic-2              	   20000	     60292 ns/op	    3380 B/op	     157 allocs/op
BenchmarkTwigGitHubAPI-2           	   20000	     90945 ns/op	    4118 B/op	     203 allocs/op
BenchmarkTwigGplusAPI-2            	  300000	      4542 ns/op	     265 B/op	      13 allocs/op
BenchmarkTwigParseAPI-2            	  200000	      8564 ns/op	     589 B/op	      26 allocs/op
BenchmarkEudoreRadixStatic-2       	   20000	     76629 ns/op	     877 B/op	       0 allocs/op
BenchmarkEudoreRadixGitHubAPI-2    	   10000	    116052 ns/op	     905 B/op	       0 allocs/op
BenchmarkEudoreRadixGplusAPI-2     	  200000	      5686 ns/op	      86 B/op	       0 allocs/op
BenchmarkEudoreRadixParseAPI-2     	  200000	     11904 ns/op	     173 B/op	       0 allocs/op
BenchmarkEudoreFullStatic-2        	   20000	     74043 ns/op	     875 B/op	       0 allocs/op
BenchmarkEudoreFullGitHubAPI-2     	   10000	    113976 ns/op	     904 B/op	       0 allocs/op
BenchmarkEudoreFullGplusAPI-2      	  200000	      5645 ns/op	      86 B/op	       0 allocs/op
BenchmarkEudoreFullParseAPI-2      	  200000	     10915 ns/op	     173 B/op	       0 allocs/op
BenchmarkHttprouterStatic-2        	2000000000	         0.00 ns/op
BenchmarkHttprouterGitHubAPI-2     	2000000000	         0.00 ns/op
BenchmarkHttprouterGplusAPI-2      	2000000000	         0.00 ns/op
BenchmarkHttprouterParseAPI-2      	2000000000	         0.00 ns/op
BenchmarkErouterRadixStatic-2      	   20000	     71185 ns/op	    2130 B/op	     157 allocs/op
BenchmarkErouterRadixGitHubAPI-2   	   20000	     96114 ns/op	    2505 B/op	     203 allocs/op
BenchmarkErouterRadixGplusAPI-2    	  300000	      4786 ns/op	     161 B/op	      13 allocs/op
BenchmarkErouterRadixParseAPI-2    	  200000	      8808 ns/op	     381 B/op	      26 allocs/op
BenchmarkErouterFullStatic-2       	   20000	     65147 ns/op	    2131 B/op	     157 allocs/op
BenchmarkErouterFullGitHubAPI-2    	   10000	    105393 ns/op	    2526 B/op	     204 allocs/op
BenchmarkErouterFullGplusAPI-2     	  300000	      4490 ns/op	     161 B/op	      13 allocs/op
BenchmarkErouterFullParseAPI-2     	  200000	      9109 ns/op	     381 B/op	      26 allocs/op
PASS
ok  	github.com/eudore/web-framework-benchmark	68.610s
```

