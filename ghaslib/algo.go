package ghaslib

import (
	"encoding/base64"
	"encoding/hex"
)

type ghas struct {
	data          []byte
	size          int
	C             chan byte
	PrintableHash func() string
}

func New(size int) *ghas {
	c := make(chan byte, 1)
	g := &ghas{size: size, C: c}
	g.PrintableHash = g.GetPrintableHex
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

func (g *ghas) Sum(dat []byte) {
	g.data = make([]byte, g.size)
	idx := 0
	for _, d := range dat {
		if idx >= g.size {
			idx = 0
		}
		g.data[idx] = hashByte(g.data[idx], d, idx)
		idx++
	}
	prev := g.data[0]
	if idx > 0 && idx < g.size {
		prev = g.data[idx-1]
	}
	for idx < g.size {
		g.data[idx] = hashByte(g.data[idx], prev, idx)
		prev = g.data[idx]
		idx++
	}
}

func (g *ghas) Eval(dat []byte) {
	g.CustomEval(dat, sendData, hashByte)
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
	prev := g.data[0]
	if idx > 0 && idx < g.size {
		prev = g.data[idx-1]
	}
	for idx < g.size {
		g.data[idx] = hashFn(g.data[idx], prev, idx)
		prev = g.data[idx]
		idx++
	}
}

func (g *ghas) GetPrintableHex() string {
	return hex.EncodeToString(g.Data())[:g.size]
}
func (g *ghas) GetPrintableB64() string {
	return base64.StdEncoding.EncodeToString(g.Data())[:g.size]
}

func sendData(g *ghas, dat []byte) {
	for _, d := range dat {
		g.C <- d
	}
	close(g.C)
}

func hashByte(prevByte byte, currentByte byte, currentIdx int) byte {
	return prevByte ^ (currentByte ^ byte(currentIdx))
}
