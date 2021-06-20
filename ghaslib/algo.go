package ghaslib

import (
	"encoding/base64"
	"encoding/hex"
)

const emptyStarter = '0'

type ghas struct {
	data          []byte
	size          int
	C             chan byte
	PrintableHash func() string
}

func New(size int) *ghas {
	c := make(chan byte, 1)
	g := &ghas{size: size, C: c}
	g.PrintableHash = g.getPrintableHex
	return g
}

func (g *ghas) Data() []byte {
	return g.data
}

func (g *ghas) String() string {
	return g.PrintableHash()
}

func (g *ghas) Size() int {
	return g.size
}

func (g *ghas) Eval(dat []byte) {
	g.CustomEval(dat, HashIt, hashByte)
}

func (g *ghas) CustomEval(dat []byte,
	hashIt func(*ghas, []byte),
	hashFn func(byte, byte, int) byte) {
	go hashIt(g, dat)

	g.data = make([]byte, g.size)
	idx := 0
	for {
		if idx >= g.size {
			idx = 0
		}
		res, ok := <-g.C
		if ok == false {
			break
		}
		g.data[idx] = hashFn(g.data[idx], res, idx)
		idx++
	}
	prev := byte(emptyStarter)
	if idx > 0 {
		prev = g.data[idx-1]
	}
	for idx < g.size {
		g.data[idx] = hashFn(g.data[idx], prev, idx)
		prev = g.data[idx]
		idx++
	}
}

func HashIt(g *ghas, dat []byte) {
	for _, d := range dat {
		g.C <- d
	}
	close(g.C)
}

func hashByte(prevByte byte, currentByte byte, currentIdx int) byte {
	return prevByte ^ (currentByte ^ byte(currentIdx))
}

func (g *ghas) getPrintableHex() string {
	return hex.EncodeToString(g.Data())[:g.size]
}
func (g *ghas) getPrintableB64() string {
	return base64.StdEncoding.EncodeToString(g.Data())[:g.size]
}
