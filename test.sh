#!/usr/bin/env bash

export GO111MODULE="on"

echo ''
go test -count=1 github.com/abhishekkr/ghas/ghaslib
echo ''
echo ''

echo ''
go test -count=1 -cover -v  github.com/abhishekkr/ghas/ghaslib
echo ''
echo ''

pushd ghaslib
echo ''
echo ''
go test -count=1 -cpuprofile=cpu.out  -bench=.
mv cpu.out ghaslib.test ../
popd
echo ''

echo ''
echo "try pprof commands like 'top10'"
go tool pprof ghaslib.test cpu.out
