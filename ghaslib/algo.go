package ghaslib

/*
No DRY applied on CustomEval & Sum; as extracting common flow here is considerably increasing exec time.
Since that is for easy maintainability and not attributing to quality of hash generated, we'll do without it.
*/

const chartable = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ:#%"
const shiftForChartable = 2 // allowing ranges (0-63) for byte<>uint8(0-255)

type ghas struct {
	data          []byte
	size          int
	C             chan byte
	PrintableHash func([]byte) string
}

func New(size int) *ghas {
	c := make(chan byte, 1)
	g := &ghas{size: size, C: c}
	g.PrintableHash = GetPrintableHash
	return g
}

func (g *ghas) Data() []byte {
	return g.data
}

func (g *ghas) String() string {
	return g.PrintableHash(g.data)
}

func (g *ghas) Size() int {
	return g.size
}

func (g *ghas) Sum(dat []byte) {
	g.data = make([]byte, g.size)
	var defVal byte = byte(len(dat)) // ensuring substring like content differs in hashes
	for i := 0; i < g.size; i++ {
		g.data[i] = defVal
	}

	idx := 0
	for _, d := range dat {
		if idx >= g.size {
			idx = 0
		}
		g.data[idx] = hashByte(g.data[idx], d, idx)
		idx++
	}
	prev := defVal
	if idx > 0 && idx < g.size {
		prev = g.data[idx-1]
	}
	for idx < g.size {
		g.data[idx] = hashByte(g.data[idx], prev, g.size-idx)
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
	var defVal byte = byte(len(dat))
	for i := 0; i < g.size; i++ {
		g.data[i] = defVal
	}

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
		g.data[idx] = hashFn(g.data[idx], prev, g.size-idx)
		prev = g.data[idx]
		idx++
	}
}

func GetPrintableHash(data []byte) string {
	hashChar := make([]byte, len(data))
	for idx, v := range data {
		hashChar[idx] = chartable[v>>shiftForChartable]
	}
	return string(hashChar)
}

func sendData(g *ghas, dat []byte) {
	for _, d := range dat {
		g.C <- d
	}
	close(g.C)
}

func hashByte(salt byte, currentByte byte, currentIdx int) byte {
	return salt ^ (currentByte ^ byte(currentIdx))
}
