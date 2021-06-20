
## Ghas

> an attempt at creating a usable Hashing Algorithm without much complexity

---

* running sample main `go run ghas.go -f $FILE_PATH -s $HASH_SIZE`

> example: `go run ghas.go -f ./LICENSE -s 256` will show time taken and hash calculated by this `Ghas` alongwith comparative calculation by Golang libs for `md5`, `sha256`, `sha512`, `hmac`
>
> sample run estimates from an avergae laptop are at [benchmarks.md](benchmarks.md)

* running unit tests, coverage, benchmark & dropping on pprof is covered by [test.sh](./test.sh)

---
