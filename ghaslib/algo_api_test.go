package ghaslib

import (
	"encoding/base64"
	"testing"
)

var (
	hash64ToDataMap = map[string]string{
		"gvgvgugugtgtgsgsgrgrgqgqgpgpgogogngngmgmglglgkgkgjgjgigighghgggg": "",
		"ononomomololokokojojoioiohohogogovovououototososororoqoqopopoooo": "a",
		"ppppononokokololoioiojojogogohohououovovososototoqoqororooooopop": "aaaaa",
		"grptjsjsjvjvjujujpjpjojojrjrjqjqjljljkjkjnjnjmjmjhjhjgjgjjjjjiji": "Ghas",
		"lusodntpsenrsvqtrrkrkqkqkpkpkokoknknkmkmklklkkkkkjkjkikikhkhkgkg": "Ghas Hash Function",
		"#PTBwHzDPMUzTOOzSSxASAMwRzDBMxOSXMDUUj#4IgUjDj#iW3HgU7KmWiX6Y0Fg": `Ghas Hash Function is just a play attempt at
	a b c d e f g h i j k l m n o p q r s t u v w x y z
	data hashing uniformly at that same length for all length of input`,
		"VKSURMRAQHIXSXtVvChYqC8UbY1w0xsLoTaVfErQnUuFmNuFaGiMlR3J3y8O2#nW": `==========================================================
What is Lorem Ipsum?
Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.

Why do we use it?
It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters, as opposed to using 'Content here, content here', making it look like readable English. Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a search for 'lorem ipsum' will uncover many web sites still in their infancy. Various versions have evolved over the years, sometimes by accident, sometimes on purpose (injected humour and the like).

Where does it come from?
Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of "de Finibus Bonorum et Malorum" (The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on the theory of ethics, very popular during the Renaissance. The first line of Lorem Ipsum, "Lorem ipsum dolor sit amet..", comes from a line in section 1.10.32.

The standard chunk of Lorem Ipsum used since the 1500s is reproduced below for those interested. Sections 1.10.32 and 1.10.33 from "de Finibus Bonorum et Malorum" by Cicero are also reproduced in their exact original form, accompanied by English versions from the 1914 translation by H. Rackham.

Where can I get some?
There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomised words which don't look even slightly believable. If you are going to use a passage of Lorem Ipsum, you need to be sure there isn't anything embarrassing hidden in the middle of text. All the Lorem Ipsum generators on the Internet tend to repeat predefined chunks as necessary, making this the first true generator on the Internet. It uses a dictionary of over 200 Latin words, combined with a handful of model sentence structures, to generate Lorem Ipsum which looks reasonable. The generated Lorem Ipsum is therefore always free from repetition, injected humour, or non-characteristic words etc.
=================================================================`,
	}

	keyToBenchmark = "VKSURMRAQHIXSXtVvChYqC8UbY1w0xsLoTaVfErQnUuFmNuFaGiMlR3J3y8O2#nW"
)

func TestAPI(t *testing.T) {
	for hash64, data := range hash64ToDataMap {
		g := New(64)
		g.Sum([]byte(data))
		ghasData := g.Data()
		ghasString := g.String()
		if len(ghasData) != 64 {
			t.Errorf("Ghas Data length doesn't match specified size.")
		}
		if len(ghasString) != 64 {
			t.Errorf("Ghas String length doesn't match specified size.")
		}
		if ghasString != hash64 {
			t.Errorf("Ghas String doesn't match expected hash:\n%v\n%v", ghasString, hash64)
		}
	}
}

func BenchmarkEval(b *testing.B) {
	baseData := []byte(hash64ToDataMap[keyToBenchmark])

	for i := 1; i < b.N; i++ {
		g := New(i)
		g.Eval(baseData)
		if len(g.String()) != i {
			b.Errorf("Ghas String getter fails in Eval.\nsize: %d\ndata: %s", i, g.String())
		}
	}
}

func BenchmarkEvalB64(b *testing.B) {
	baseData := []byte(hash64ToDataMap[keyToBenchmark])

	for i := 1; i < b.N; i++ {
		g := New(i)
		g.PrintableHash = GetPrintableB64
		g.Eval(baseData)
		if len(g.String()) != i {
			b.Errorf("Ghas String getter fails in Eval. %s", g.String())
		}
	}
}

func BenchmarkSum(b *testing.B) {
	baseData := []byte(hash64ToDataMap[keyToBenchmark])

	for i := 1; i < b.N; i++ {
		g := New(i)
		g.Sum(baseData)
		if len(g.String()) != i {
			b.Errorf("Ghas String getter fails in Eval. %s", g.String())
		}
	}
}

func BenchmarkSumB64(b *testing.B) {
	baseData := []byte(hash64ToDataMap[keyToBenchmark])

	for i := 1; i < b.N; i++ {
		g := New(i)
		g.PrintableHash = GetPrintableB64
		g.Sum(baseData)
		if len(g.String()) != i {
			b.Errorf("Ghas String getter fails in Eval. %s", g.String())
		}
	}
}

func GetPrintableB64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)[:len(data)]
}
