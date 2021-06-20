
## Summary

* Public function `Sum` runs faster due to lack of locks from goroutine implementation `Eval`

* As of simple internal clock measures it's; while producing `256 char long checksum`

> * slower than MD5 (fastest), SHA256, SHA512, HMAC for data smaller than 20KB (although stays mostly same in speed when producing similar char long output)
>
> * starts being useful/faster for data/files larger than 20KB even with higher char count checksum
>
> * increasing or decreasing checksum size doesn't noticeably impact the perf

---

## Test Results

```
# ./test.sh

ok  	github.com/abhishekkr/ghas/ghaslib	0.001s

...
coverage: 97.8% of statements
ok  	github.com/abhishekkr/ghas/ghaslib	0.001s	coverage: 97.8% of statements

goos: linux
goarch: amd64
pkg: github.com/abhishekkr/ghas/ghaslib
BenchmarkEvalHex-8   	    2742	    399056 ns/op
BenchmarkEvalB64-8   	    3028	    392148 ns/op
BenchmarkSumHex-8    	   77977	    121189 ns/op
BenchmarkSumB64-8    	   84766	    117913 ns/op
PASS
ok  	github.com/abhishekkr/ghas/ghaslib	22.280s

File: ghaslib.test
Type: cpu
...
Showing top 5 nodes out of 107
      flat  flat%   sum%        cum   cum%
    9650ms 33.98% 33.98%    10560ms 37.18%  github.com/abhishekkr/ghas/ghaslib.(*ghas).Sum
    3390ms 11.94% 45.92%     3440ms 12.11%  encoding/hex.Encode
    3330ms 11.73% 57.64%     3330ms 11.73%  encoding/base64.(*Encoding).Encode
    1720ms  6.06% 63.70%     2710ms  9.54%  runtime.scanobject
     750ms  2.64% 66.34%      750ms  2.64%  runtime.memclrNoHeapPointers
```

---

## Comparative Runs

* for a 43 char containing regular file `go.mod`; slower than rest

```
# go run ghas.go -f go.mod -s 256
file length:  43

-GHAS-> 6.716µs

-MD5-> 589ns

-SHA256-> 1.45µs

-SHA512-> 1.491µs

-HMAC512-> 5.416µs
```

* for a 28KB pprof dump

```
# go run ghas.go -f cpu.out -s 256
file length:  24876

-GHAS-> 31.288µs

-MD5-> 35.808µs

-SHA256-> 67.566µs

-SHA512-> 52.136µs

-HMAC512-> 50.429µs
```


* for a 2.1MB PDF file

```
# go run ghas.go -f 1706.pdf -s 256
file length:  2128686

-GHAS-> 1.996587ms

-MD5-> 3.004071ms

-SHA256-> 5.324719ms

-SHA512-> 3.604116ms

-HMAC512-> 3.629344ms
```

* for 696MB ISO file

```
# go run ghas.go -f ~/Desktop/archlinux-2021.02.01-x86_64.iso -s 256
file length:  729100288

-GHAS-> 625.884958ms

-MD5-> 937.482085ms

-SHA256-> 1.591005546s

-SHA512-> 1.077741687s

-HMAC512-> 1.079296026s
```

---
