
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
BenchmarkEval-8      	    1730	    596685 ns/op
BenchmarkEvalB64-8   	    1987	    586590 ns/op
BenchmarkSum-8       	   47719	    106051 ns/op
BenchmarkSumB64-8    	   48456	    118709 ns/op
PASS
ok  	github.com/abhishekkr/ghas/ghaslib	13.756s
~/ABK/dev/abhishekkr/on_github/ghas


try pprof commands like 'top10'
File: ghaslib.test
Type: cpu
Time: Jun 23, 2021 at 3:35pm (IST)
Duration: 13.75s, Total samples = 15.87s (115.43%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top5
Showing nodes accounting for 10770ms, 67.86% of 15870ms total
Dropped 144 nodes (cum <= 79.35ms)
Showing top 5 nodes out of 106
      flat  flat%   sum%        cum   cum%
    6640ms 41.84% 41.84%     7690ms 48.46%  github.com/abhishekkr/ghas/ghaslib.(*ghas).Sum
    1770ms 11.15% 52.99%     1770ms 11.15%  encoding/base64.(*Encoding).Encode
     990ms  6.24% 59.23%     1410ms  8.88%  github.com/abhishekkr/ghas/ghaslib.GetPrintableHash
     760ms  4.79% 64.02%      760ms  4.79%  github.com/abhishekkr/ghas/ghaslib.hashByte
     610ms  3.84% 67.86%      880ms  5.55%  runtime.scanobject
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
