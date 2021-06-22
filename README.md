
## Ghas

> an attempt at creating a usable Hashing Algorithm

---

* running sample main `go run ghas.go -f $FILE_PATH -s $HASH_SIZE`

```
Usage of ghas:
  -v    verbose mode, to enable file length & time taken info for GHAS
  -c    prints comparative hash from MD5, SHA256, SHA512, HMAC with time taken (enables verbose by default)
  -f string
        path of file to hash
  -s int
        size of hash to generate (default 64)
```

> example: `go run ghas.go -f ./LICENSE -s 256 -c` will show time taken and hash calculated by this `Ghas` alongwith comparative calculation by Golang libs for `md5`, `sha256`, `sha512`, `hmac`
>
> sample run estimates from an avergae laptop are at [benchmarks.md](benchmarks.md)

* running unit tests, coverage, benchmark & dropping on pprof is covered by [test.sh](./test.sh)

---
