# Web Framework Benchmark

## Installation

- `go get github.com/eudore/web-framework-benchmark`

## Running Benchmark

```sh
go test -bench=. github.com/eudore/web-framework-benchmark
```

环境变量NUM是加载的空中间件数量，数据为空就是没有增加加载中间件。

## Resule

时间: 2020年2月27日

基于echo使用的[测试方法](https://github.com/vishr/web-framework-benchmark)fork ，

iris无法运行，会panic，测试代码已注释。

```
[root@izj6cffbpd9lzl3tcm2csxz eudore]# go version
go version go1.13 linux/amd64
[root@izj6cffbpd9lzl3tcm2csxz web-framework-benchmark]# go test -bench=. github.com/eudore/web-framework-benchmark
goos: linux
goarch: amd64
pkg: github.com/eudore/web-framework-benchmark
BenchmarkGolfStatic-2              	   30000	     42354 ns/op	    2432 B/op	     157 allocs/op
BenchmarkGolfGitHubAPI-2           	   20000	     59488 ns/op	    2526 B/op	     203 allocs/op
BenchmarkGolfGplusAPI-2            	  500000	      2782 ns/op	     173 B/op	      13 allocs/op
BenchmarkGolfParseAPI-2            	  300000	      4606 ns/op	     323 B/op	      26 allocs/op
BenchmarkEchoStatic-2              	   30000	     54740 ns/op	    2413 B/op	     157 allocs/op
BenchmarkEchoGitHubAPI-2           	   20000	     76689 ns/op	    2496 B/op	     203 allocs/op
BenchmarkEchoGplusAPI-2            	  300000	      4178 ns/op	     161 B/op	      13 allocs/op
BenchmarkEchoParseAPI-2            	  200000	      7647 ns/op	     381 B/op	      26 allocs/op
BenchmarkGinStatic-2               	   20000	     79573 ns/op	    8405 B/op	     157 allocs/op
BenchmarkGinGitHubAPI-2            	   10000	    116427 ns/op	   10620 B/op	     203 allocs/op
BenchmarkGinGplusAPI-2             	  200000	      6501 ns/op	     710 B/op	      13 allocs/op
BenchmarkGinParseAPI-2             	  100000	     11981 ns/op	    1421 B/op	      26 allocs/op
BenchmarkDotwebStatic-2            	    5000	    357222 ns/op	  101237 B/op	    2847 allocs/op
BenchmarkDotwebGitHubAPI-2         	    2000	    518961 ns/op	  145389 B/op	    3856 allocs/op
BenchmarkDotwebGplusAPI-2          	   50000	     29051 ns/op	    9042 B/op	     246 allocs/op
BenchmarkDotwebParseAPI-2          	   30000	     56751 ns/op	   17454 B/op	     488 allocs/op
BenchmarkBeegoStatic-2             	    5000	    260340 ns/op	   77517 B/op	     942 allocs/op
BenchmarkBeegoGitHubAPI-2          	    3000	    356271 ns/op	  104920 B/op	    1222 allocs/op
BenchmarkBeegoGplusAPI-2           	  100000	     20012 ns/op	    6437 B/op	      78 allocs/op
BenchmarkBeegoParseAPI-2           	   50000	     41463 ns/op	   12877 B/op	     156 allocs/op
BenchmarkTwigStatic-2              	   20000	     59363 ns/op	    3380 B/op	     157 allocs/op
BenchmarkTwigGitHubAPI-2           	   20000	     90016 ns/op	    4118 B/op	     203 allocs/op
BenchmarkTwigGplusAPI-2            	  300000	      4980 ns/op	     265 B/op	      13 allocs/op
BenchmarkTwigParseAPI-2            	  200000	      8679 ns/op	     589 B/op	      26 allocs/op
BenchmarkEudoreRadixStatic-2       	   30000	     58397 ns/op	    2417 B/op	     157 allocs/op
BenchmarkEudoreRadixGitHubAPI-2    	   20000	     90214 ns/op	    2509 B/op	     203 allocs/op
BenchmarkEudoreRadixGplusAPI-2     	  300000	      4257 ns/op	     161 B/op	      13 allocs/op
BenchmarkEudoreRadixParseAPI-2     	  200000	      8074 ns/op	     381 B/op	      26 allocs/op
BenchmarkEudoreFullStatic-2        	   30000	     57943 ns/op	    2416 B/op	     157 allocs/op
BenchmarkEudoreFullGitHubAPI-2     	   20000	     91842 ns/op	    2509 B/op	     203 allocs/op
BenchmarkEudoreFullGplusAPI-2      	  300000	      4472 ns/op	     161 B/op	      13 allocs/op
BenchmarkEudoreFullParseAPI-2      	  200000	      8294 ns/op	     381 B/op	      26 allocs/op
BenchmarkHttprouterStatic-2        	   50000	     25276 ns/op	    1949 B/op	     157 allocs/op
BenchmarkHttprouterGitHubAPI-2     	   30000	     54370 ns/op	   16571 B/op	     370 allocs/op
BenchmarkHttprouterGplusAPI-2      	  500000	      2547 ns/op	     813 B/op	      24 allocs/op
BenchmarkHttprouterParseAPI-2      	  300000	      3887 ns/op	     963 B/op	      42 allocs/op
BenchmarkErouterRadixStatic-2      	   50000	     35762 ns/op	    1950 B/op	     157 allocs/op
BenchmarkErouterRadixGitHubAPI-2   	   30000	     59412 ns/op	    2786 B/op	     203 allocs/op
BenchmarkErouterRadixGplusAPI-2    	  500000	      2593 ns/op	     173 B/op	      13 allocs/op
BenchmarkErouterRadixParseAPI-2    	  300000	      4372 ns/op	     323 B/op	      26 allocs/op
BenchmarkErouterFullStatic-2       	   50000	     35829 ns/op	    1950 B/op	     157 allocs/op
BenchmarkErouterFullGitHubAPI-2    	   20000	     60026 ns/op	    2504 B/op	     203 allocs/op
BenchmarkErouterFullGplusAPI-2     	  500000	      2573 ns/op	     173 B/op	      13 allocs/op
BenchmarkErouterFullParseAPI-2     	  300000	      4636 ns/op	     323 B/op	      26 allocs/op
PASS
ok  	github.com/eudore/web-framework-benchmark	69.344s
```

```
[root@izj6cffbpd9lzl3tcm2csxz web-framework-benchmark]# go version
go version go1.14 linux/amd64
[root@izj6cffbpd9lzl3tcm2csxz web-framework-benchmark]# go test -bench=. github.com/eudore/web-framework-benchmark
goos: linux
goarch: amd64
pkg: github.com/eudore/web-framework-benchmark
BenchmarkGolfStatic-2              	   30081	     39556 ns/op	    2429 B/op	     157 allocs/op
BenchmarkGolfGitHubAPI-2           	   20692	     59314 ns/op	    2496 B/op	     203 allocs/op
BenchmarkGolfGplusAPI-2            	  470106	      2653 ns/op	     177 B/op	      13 allocs/op
BenchmarkGolfParseAPI-2            	  257791	      4494 ns/op	     342 B/op	      26 allocs/op
BenchmarkEchoStatic-2              	   24610	     47909 ns/op	    1963 B/op	     157 allocs/op
BenchmarkEchoGitHubAPI-2           	   15925	     73091 ns/op	    2719 B/op	     203 allocs/op
BenchmarkEchoGplusAPI-2            	  321218	      3876 ns/op	     157 B/op	      13 allocs/op
BenchmarkEchoParseAPI-2            	  173895	      7041 ns/op	     407 B/op	      26 allocs/op
BenchmarkGinStatic-2               	   15340	     79017 ns/op	    8669 B/op	     157 allocs/op
BenchmarkGinGitHubAPI-2            	   10000	    113548 ns/op	   10620 B/op	     203 allocs/op
BenchmarkGinGplusAPI-2             	  192846	      6511 ns/op	     713 B/op	      13 allocs/op
BenchmarkGinParseAPI-2             	  101425	     12079 ns/op	    1419 B/op	      26 allocs/op
BenchmarkDotwebStatic-2            	    3729	    324080 ns/op	  100645 B/op	    2835 allocs/op
BenchmarkDotwebGitHubAPI-2         	    2504	    461417 ns/op	  143386 B/op	    3834 allocs/op
BenchmarkDotwebGplusAPI-2          	   44096	     27226 ns/op	    8976 B/op	     245 allocs/op
BenchmarkDotwebParseAPI-2          	   21306	     54921 ns/op	   17321 B/op	     485 allocs/op
BenchmarkBeegoStatic-2             	    5402	    228641 ns/op	   76511 B/op	    1099 allocs/op
BenchmarkBeegoGitHubAPI-2          	    3655	    323591 ns/op	   99193 B/op	    1423 allocs/op
BenchmarkBeegoGplusAPI-2           	   67594	     17667 ns/op	    6332 B/op	      91 allocs/op
BenchmarkBeegoParseAPI-2           	   36033	     33862 ns/op	   12656 B/op	     182 allocs/op
BenchmarkTwigStatic-2              	   21424	     59551 ns/op	    3323 B/op	     157 allocs/op
BenchmarkTwigGitHubAPI-2           	   14194	     82770 ns/op	    4473 B/op	     203 allocs/op
BenchmarkTwigGplusAPI-2            	  270946	      4431 ns/op	     271 B/op	      13 allocs/op
BenchmarkTwigParseAPI-2            	  137424	      8018 ns/op	     542 B/op	      26 allocs/op
BenchmarkEudoreRadixStatic-2       	   19804	     58669 ns/op	    2141 B/op	     157 allocs/op
BenchmarkEudoreRadixGitHubAPI-2    	   13033	     91678 ns/op	    2983 B/op	     203 allocs/op
BenchmarkEudoreRadixGplusAPI-2     	  273674	      4387 ns/op	     167 B/op	      13 allocs/op
BenchmarkEudoreRadixParseAPI-2     	  141092	      8023 ns/op	     331 B/op	      26 allocs/op
BenchmarkEudoreFullStatic-2        	   21002	     57398 ns/op	    2090 B/op	     157 allocs/op
BenchmarkEudoreFullGitHubAPI-2     	   12312	     95009 ns/op	    3063 B/op	     203 allocs/op
BenchmarkEudoreFullGplusAPI-2      	  281281	      4540 ns/op	     165 B/op	      13 allocs/op
BenchmarkEudoreFullParseAPI-2      	  139195	      8289 ns/op	     332 B/op	      26 allocs/op
BenchmarkHttprouterStatic-2        	   61990	     20023 ns/op	    1431 B/op	     157 allocs/op
BenchmarkHttprouterGitHubAPI-2     	   21801	     51615 ns/op	   15788 B/op	     370 allocs/op
BenchmarkHttprouterGplusAPI-2      	  446752	      2366 ns/op	     743 B/op	      24 allocs/op
BenchmarkHttprouterParseAPI-2      	  315570	      3605 ns/op	     801 B/op	      42 allocs/op
BenchmarkErouterRadixStatic-2      	   38384	     30774 ns/op	    1218 B/op	     157 allocs/op
BenchmarkErouterRadixGitHubAPI-2   	   23169	     52165 ns/op	    1910 B/op	     203 allocs/op
BenchmarkErouterRadixGplusAPI-2    	  578235	      2116 ns/op	      85 B/op	      13 allocs/op
BenchmarkErouterRadixParseAPI-2    	  316267	      3858 ns/op	     161 B/op	      26 allocs/op
BenchmarkErouterFullStatic-2       	   37245	     31084 ns/op	    1246 B/op	     157 allocs/op
BenchmarkErouterFullGitHubAPI-2    	   22321	     52936 ns/op	    1969 B/op	     203 allocs/op
BenchmarkErouterFullGplusAPI-2     	  568222	      2189 ns/op	      86 B/op	      13 allocs/op
BenchmarkErouterFullParseAPI-2     	  322863	      3743 ns/op	     159 B/op	      26 allocs/op
PASS
ok  	github.com/eudore/web-framework-benchmark	65.149s
```
