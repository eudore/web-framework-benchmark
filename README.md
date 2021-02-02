# Web Framework Benchmark

golf、echo、gin、beego、iris(panic无法执行)、twig、eudore七框架和httprouter、chi、eudore-router三纯路由器进行路由匹配性能测试。

环境变量NUM是加载的空中间件数量，数据为空就是没有增加加载中间件，仅echo框架中间件运行时闭包消耗性能严重。

## Installation

- `go get github.com/eudore/web-framework-benchmark`

## Running Benchmark

```sh
go test -bench=. github.com/eudore/web-framework-benchmark
```

## Result

时间: 2021年2月2日

基于echo使用的[测试方法](https://github.com/vishr/web-framework-benchmark)fork ，

iris v12无法运行单元测试会触发panic，测试代码已注释。

```
[root@izuf6b7o9hu1q8vzyvkyznz web-framework-benchmark]# go version
go version go1.15.3 linux/amd64
[root@izuf6b7o9hu1q8vzyvkyznz web-framework-benchmark]# GO111MODULE=on go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/eudore/web-framework-benchmark
BenchmarkGolfStatic          	   28286	     42739 ns/op	    2503 B/op	     157 allocs/op
BenchmarkGolfGitHubAPI       	   18978	     60247 ns/op	    2574 B/op	     203 allocs/op
BenchmarkGolfGplusAPI        	  441594	      2831 ns/op	     182 B/op	      13 allocs/op
BenchmarkGolfParseAPI        	  255552	      4726 ns/op	     343 B/op	      26 allocs/op
BenchmarkEchoStatic          	   21883	     51482 ns/op	    2051 B/op	     157 allocs/op
BenchmarkEchoGitHubAPI       	   15219	     79214 ns/op	    2770 B/op	     203 allocs/op
BenchmarkEchoGplusAPI        	  295573	      4232 ns/op	     162 B/op	      13 allocs/op
BenchmarkEchoParseAPI        	  170500	      7565 ns/op	     411 B/op	      26 allocs/op
BenchmarkGinStatic           	   17946	     63500 ns/op	    8504 B/op	     157 allocs/op
BenchmarkGinGitHubAPI        	   12592	     91273 ns/op	   11126 B/op	     203 allocs/op
BenchmarkGinGplusAPI         	  222997	      5182 ns/op	     701 B/op	      13 allocs/op
BenchmarkGinParseAPI         	  114259	      9738 ns/op	    1399 B/op	      26 allocs/op
BenchmarkBeegoStatic         	    4198	    300368 ns/op	   76741 B/op	    1099 allocs/op
BenchmarkBeegoGitHubAPI      	    2805	    444280 ns/op	   99590 B/op	    1424 allocs/op
BenchmarkBeegoGplusAPI       	   54568	     22940 ns/op	    6347 B/op	      91 allocs/op
BenchmarkBeegoParseAPI       	   28390	     43347 ns/op	   12688 B/op	     182 allocs/op
BenchmarkTwigStatic          	   19267	     60687 ns/op	    3413 B/op	     157 allocs/op
BenchmarkTwigGitHubAPI       	   13252	     88902 ns/op	    4560 B/op	     203 allocs/op
BenchmarkTwigGplusAPI        	  266600	      4834 ns/op	     273 B/op	      13 allocs/op
BenchmarkTwigParseAPI        	  128240	      8920 ns/op	     551 B/op	      26 allocs/op
BenchmarkEudoreStatic        	   21118	     55109 ns/op	    2096 B/op	     157 allocs/op
BenchmarkEudoreGitHubAPI     	   13137	     87902 ns/op	    3013 B/op	     204 allocs/op
BenchmarkEudoreGplusAPI      	  330379	      3981 ns/op	     156 B/op	      13 allocs/op
BenchmarkEudoreParseAPI      	  177199	      7297 ns/op	     403 B/op	      26 allocs/op
BenchmarkHttprouterStatic    	   55825	     21407 ns/op	    1554 B/op	     157 allocs/op
BenchmarkHttprouterGitHubAPI 	   21518	     52940 ns/op	   15809 B/op	     370 allocs/op
BenchmarkHttprouterGplusAPI  	  458434	      2523 ns/op	     741 B/op	      24 allocs/op
BenchmarkHttprouterParseAPI  	  293004	      3910 ns/op	     810 B/op	      42 allocs/op
BenchmarkChiStatic           	    9190	    139410 ns/op	   69094 B/op	     628 allocs/op
BenchmarkChiGitHubAPI        	    6774	    189363 ns/op	   89398 B/op	     812 allocs/op
BenchmarkChiGplusAPI         	  112657	      9749 ns/op	    5719 B/op	      52 allocs/op
BenchmarkChiParseAPI         	   62676	     18495 ns/op	   11423 B/op	     104 allocs/op
BenchmarkErouterStatic       	   34785	     33387 ns/op	    1312 B/op	     157 allocs/op
BenchmarkErouterGitHubAPI    	   21375	     56848 ns/op	    2038 B/op	     203 allocs/op
BenchmarkErouterGplusAPI     	  545409	      2328 ns/op	      89 B/op	      13 allocs/op
BenchmarkErouterParseAPI     	  302883	      4098 ns/op	     166 B/op	      26 allocs/op
PASS
ok  	github.com/eudore/web-framework-benchmark	54.048s
```
