package ghaslib

import (
	"bytes"
	"reflect"
	"strconv"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	size := 10
	g := New(size)
	if g.size != size {
		t.Errorf("Ghas size doesn't apply as passed.")
	}
	if reflect.TypeOf(g.C).String() != "chan uint8" {
		t.Errorf("Ghas channel type doesn't match.")
	}
	if reflect.TypeOf(g).String() != "*ghaslib.ghas" {
		t.Errorf("Ghas type doesn't match.")
	}
}

func TestGhasData(t *testing.T) {
	data := []byte("ghas")
	g := &ghas{data: data}
	if !bytes.Equal(g.Data(), data) {
		t.Errorf("Ghas Data getter fails. %v", g.Data())
	}
}

func TestGhasString(t *testing.T) {
	g := &ghas{data: []byte("ghas"), size: 4}
	g.PrintableHash = g.getPrintableB64
	if g.String() != "Z2hh" {
		t.Errorf("Ghas String getter fails. %s", g.String())
	}
}

func TestGhasSize(t *testing.T) {
	size := 8
	g := &ghas{size: size}
	if g.Size() != size {
		t.Errorf("Ghas Size getter fails.")
	}
}

func TestGhasEval(t *testing.T) {
	str := "ghas"
	g := New(8)
	g.Eval([]byte(str))
	if !bytes.Equal(g.data, []byte{103, 105, 99, 112, 116, 113, 119, 112}) {
		t.Errorf("Ghas has wrong data in Eval. %v", g.data)
	}
	if g.String() != "67696370" {
		t.Errorf("Ghas String getter fails in Eval. %s", g.String())
	}
}

func TestGhasCustomEval(t *testing.T) {
	str := "ghas"
	g := New(8)
	zeroHashIt := func(g *ghas, b []byte) {
		for idx, d := range b {
			g.C <- (d + []byte(strconv.Itoa(idx))[0])
		}
		close(g.C)
	}

	hashByte := func(prevByte byte, currentByte byte, currentIdx int) byte {
		return currentByte
	}
	g.CustomEval([]byte(str), zeroHashIt, hashByte)
	if !bytes.Equal(g.data, []byte{151, 153, 147, 166, 166, 166, 166, 166}) {
		t.Errorf("Ghas has wrong data in CustomEval. %v", g.data)
	}
}

func TestHashIt(t *testing.T) {
	var expectedByte byte = 101
	g := &ghas{size: 1, C: make(chan byte, 1)}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		b, ok := <-g.C
		if !ok || b != expectedByte {
			t.Errorf("Ghas HashIt doesn't send on channel.")
		}
	}()
	HashIt(g, []byte{expectedByte})
	wg.Wait()
}
