package ghaslib

import (
	"bytes"
	"reflect"
	"strconv"
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
	g.PrintableHash = g.GetPrintableB64
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

func TestGhasSum(t *testing.T) {
	str := "ghas"
	g := New(8)
	g.Sum([]byte(str))
	if g.String() != "636d6774" {
		t.Errorf("Ghas String getter fails in Eval. %s", g.String())
	}
}

func TestGhasEval(t *testing.T) {
	str := "ghas"
	g := New(8)
	g.Eval([]byte(str))
	if g.String() != "636d6774" {
		t.Errorf("Ghas String getter fails in Eval. %s", g.String())
	}
}

func TestGhasCustomEval(t *testing.T) {
	str := "ghas"
	g := New(8)
	sendZero := func(g *ghas, b []byte) {
		for idx, d := range b {
			g.C <- (d + []byte(strconv.Itoa(idx))[0])
		}
		close(g.C)
	}

	hashByte := func(salt byte, currentByte byte, currentIdx int) byte {
		return currentByte
	}
	g.CustomEval([]byte(str), sendZero, hashByte)
	if !bytes.Equal(g.data, []byte{151, 153, 147, 166, 166, 166, 166, 166}) {
		t.Errorf("Ghas has wrong data in CustomEval. %v", g.data)
	}
}
