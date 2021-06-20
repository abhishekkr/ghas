package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"time"

	ghaslib "github.com/abhishekkr/ghas/ghaslib"

	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

func main() {
	comparative := flag.Bool("c", false, "prints comparative hash from MD5, SHA256, SHA512, HMAC with time taken")
	prnToHash := flag.String("p", "hex", "uses 'hex' converter by default for byte to string, can use 'b64' for base64")
	sizeToHash := flag.Int("s", 64, "size of hash to generate")
	fileToHash := flag.String("f", "", "path of file to hash")
	flag.Parse()
	if *fileToHash != "" {
		hashTheFile(*fileToHash, *sizeToHash, *prnToHash, *comparative)
	}
}

func checkFatal(e error) {
	if e != nil {
		panic(e)
	}
}

func hashTheFile(filepath string, hashSize int, prnToHash string, comparative bool) {
	dat, err := ioutil.ReadFile(filepath)
	fmt.Println("file length: ", len(dat))
	checkFatal(err)
	start := time.Now()
	ghash := GetGhas(dat, hashSize, prnToHash)
	fmt.Println("\n-GHAS->", time.Since(start), " | for hash with", len(ghash), "bytes")
	fmt.Println(ghash)
	if comparative {
		otherHashing(dat, prnToHash)
	}
}

func GetGhas(dat []byte, hashSize int, prnToHash string) string {
	g := ghaslib.New(hashSize)
	if prnToHash == "b64" {
		g.PrintableHash = g.GetPrintableB64
	}
	//g.Eval([]byte(dat))
	g.Sum([]byte(dat))
	return g.String()
}

func otherHashing(dat []byte, prnToHash string) {
	mstart := time.Now()
	mhashB := md5.Sum(dat)
	mhash := hex.EncodeToString(mhashB[:])
	fmt.Println("\n-MD5->", time.Since(mstart), " | for hash with", len(mhash), "bytes")
	fmt.Println(mhash)

	s2start := time.Now()
	s256B := sha256.Sum256(dat)
	s256 := hex.EncodeToString(s256B[:])
	fmt.Println("\n-SHA256->", time.Since(s2start), " | for hash with", len(s256), "bytes")
	fmt.Println(s256)

	s5start := time.Now()
	s512B := sha512.Sum512(dat)
	s512 := hex.EncodeToString(s512B[:])
	fmt.Println("\n-SHA512->", time.Since(s5start), " | for hash with", len(s512), "bytes")
	fmt.Println(s512)

	hstart := time.Now()
	hmac512 := hmac.New(sha512.New, []byte("secret"))
	hmac512.Write(dat)
	h512 := base64.StdEncoding.EncodeToString(hmac512.Sum(nil))
	fmt.Println("\n-HMAC512->", time.Since(hstart), " | for hash with", len(h512), "bytes")
	fmt.Println(h512)
}
