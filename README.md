# Web Framework Benchmark

## Installation

- `go get github.com/eudore/web-framework-benchmark`

## Running Benchmark

```sh
go test -bench=. github.com/eudore/web-framework-benchmark
```

环境变量NUM是加载的空中间件数量，数据为空就是没有增加加载中间件。

## Resule

时间: 2019年10月27日

基于echo使用的[测试方法](https://github.com/vishr/web-framework-benchmark)fork ，

iris无法运行，会panic，测试代码已注释。

```
[root@izj6cffbpd9lzl3tcm2csxz eudore]# go version
go version go1.13 linux/amd64
[root@izj6cffbpd9lzl3tcm2csxz web-framework-benchmark]# go test -bench=. github.com/eudore/web-framework-benchmark
goos: linux
goarch: amd64
pkg: github.com/eudore/web-framework-benchmark
BenchmarkGolfStatic-2              	   30208	     44498 ns/op	    2424 B/op	     157 allocs/op
BenchmarkGolfGitHubAPI-2           	   17150	     72195 ns/op	    2676 B/op	     204 allocs/op
BenchmarkGolfGplusAPI-2            	  406230	      2833 ns/op	     189 B/op	      13 allocs/op
BenchmarkGolfParseAPI-2            	  235088	      5087 ns/op	     355 B/op	      26 allocs/op
BenchmarkEchoStatic-2              	   22947	     53032 ns/op	    2015 B/op	     157 allocs/op
BenchmarkEchoGitHubAPI-2           	   16620	     71640 ns/op	    2673 B/op	     203 allocs/op
BenchmarkEchoGplusAPI-2            	  307123	      3983 ns/op	     160 B/op	      13 allocs/op
BenchmarkEchoParseAPI-2            	  174658	      7029 ns/op	     406 B/op	      26 allocs/op
BenchmarkGinStatic-2               	   14479	     78531 ns/op	    8736 B/op	     157 allocs/op
BenchmarkGinGitHubAPI-2            	   10000	    114445 ns/op	   10620 B/op	     203 allocs/op
BenchmarkGinGplusAPI-2             	  189676	      6625 ns/op	     715 B/op	      13 allocs/op
BenchmarkGinParseAPI-2             	  102295	     12541 ns/op	    1417 B/op	      26 allocs/op
BenchmarkDotwebStatic-2            	    3458	    366361 ns/op	  101343 B/op	    2838 allocs/op
BenchmarkDotwebGitHubAPI-2         	    2470	    551926 ns/op	  144047 B/op	    3837 allocs/op
BenchmarkDotwebGplusAPI-2          	   37876	     35167 ns/op	    8981 B/op	     246 allocs/op
BenchmarkDotwebParseAPI-2          	   17848	     62844 ns/op	   17320 B/op	     485 allocs/op
BenchmarkBeegoStatic-2             	    4576	    257102 ns/op	   76660 B/op	    1099 allocs/op
BenchmarkBeegoGitHubAPI-2          	    3243	    368925 ns/op	   99363 B/op	    1423 allocs/op
BenchmarkBeegoGplusAPI-2           	   57676	     21200 ns/op	    6343 B/op	      91 allocs/op
BenchmarkBeegoParseAPI-2           	   28618	     39455 ns/op	   12688 B/op	     182 allocs/op
BenchmarkTwigStatic-2              	   20112	     59085 ns/op	    3376 B/op	     157 allocs/op
BenchmarkTwigGitHubAPI-2           	   13903	     85844 ns/op	    4499 B/op	     203 allocs/op
BenchmarkTwigGplusAPI-2            	  267199	      5004 ns/op	     272 B/op	      13 allocs/op
BenchmarkTwigParseAPI-2            	  143353	      8440 ns/op	     536 B/op	      26 allocs/op
BenchmarkEudoreRadixStatic-2       	   22930	     61829 ns/op	    2018 B/op	     157 allocs/op
BenchmarkEudoreRadixGitHubAPI-2    	   13875	     83999 ns/op	    2895 B/op	     203 allocs/op
BenchmarkEudoreRadixGplusAPI-2     	  289152	      4076 ns/op	     164 B/op	      13 allocs/op
BenchmarkEudoreRadixParseAPI-2     	  134761	      7574 ns/op	     336 B/op	      26 allocs/op
BenchmarkEudoreFullStatic-2        	   22110	     55446 ns/op	    2046 B/op	     157 allocs/op
BenchmarkEudoreFullGitHubAPI-2     	   13608	     90673 ns/op	    2920 B/op	     203 allocs/op
BenchmarkEudoreFullGplusAPI-2      	  261680	      4185 ns/op	     170 B/op	      13 allocs/op
BenchmarkEudoreFullParseAPI-2      	  160928	      7761 ns/op	     315 B/op	      26 allocs/op
BenchmarkHttprouterStatic-2        	   58454	     20522 ns/op	    1498 B/op	     157 allocs/op
BenchmarkHttprouterGitHubAPI-2     	   23164	     51439 ns/op	   15694 B/op	     370 allocs/op
BenchmarkHttprouterGplusAPI-2      	  525870	      2425 ns/op	     731 B/op	      24 allocs/op
BenchmarkHttprouterParseAPI-2      	  338446	      3789 ns/op	     896 B/op	      42 allocs/op
BenchmarkErouterRadixStatic-2      	   37615	     32547 ns/op	    1237 B/op	     157 allocs/op
BenchmarkErouterRadixGitHubAPI-2   	   23498	     54143 ns/op	    1889 B/op	     203 allocs/op
BenchmarkErouterRadixGplusAPI-2    	  521916	      2263 ns/op	      92 B/op	      13 allocs/op
BenchmarkErouterRadixParseAPI-2    	  299391	      3961 ns/op	     167 B/op	      26 allocs/op
BenchmarkErouterFullStatic-2       	   33639	     31997 ns/op	    1346 B/op	     157 allocs/op
BenchmarkErouterFullGitHubAPI-2    	   21007	     58242 ns/op	    1243 B/op	     203 allocs/op
BenchmarkErouterFullGplusAPI-2     	  553612	      2199 ns/op	      88 B/op	      13 allocs/op
BenchmarkErouterFullParseAPI-2     	  297276	      4070 ns/op	     168 B/op	      26 allocs/op
PASS
ok  	github.com/eudore/web-framework-benchmark	69.344s
```

