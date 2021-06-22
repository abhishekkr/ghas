
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
# ./ghas -f go.mod -s 256
file length:  43

-GHAS-> 4.044µs  | for hash with 256 bytes

-MD5-> 836ns  | for hash with 32 bytes

-SHA256-> 1.77µs  | for hash with 64 bytes

-SHA512-> 1.315µs  | for hash with 128 bytes

-HMAC512-> 3.28µs  | for hash with 88 bytes
```

* for a 20KB `cpu.out` dump

```
# ./ghas -f cpu.out -s 256
file length:  20178

-GHAS-> 25.557µs  | for hash with 256 bytes

-MD5-> 31.673µs  | for hash with 32 bytes

-SHA256-> 56.55µs  | for hash with 64 bytes

-SHA512-> 39.427µs  | for hash with 128 bytes

-HMAC512-> 40.123µs  | for hash with 88 bytes
```


* for a 2.1MB PDF file

```
# ./ghas -f 1706.pdf -s 256 -c
file length:  2128686

-GHAS-> 2.411464ms  | for hash with 256 bytes

-MD5-> 3.247847ms  | for hash with 32 bytes

-SHA256-> 5.76674ms  | for hash with 64 bytes

-SHA512-> 4.429595ms  | for hash with 128 bytes

-HMAC512-> 4.027916ms  | for hash with 88 bytes
```

* for 696MB ISO file

```
# ./ghas -f archlinux-2021.02.01-x86_64.iso -s 256 -c
file length:  729100288

-GHAS-> 730.098234ms  | for hash with 256 bytes

-MD5-> 977.162613ms  | for hash with 32 bytes

-SHA256-> 1.660822319s  | for hash with 64 bytes

-SHA512-> 1.134428145s  | for hash with 128 bytes

-HMAC512-> 1.126056971s  | for hash with 88 bytes
```

---
