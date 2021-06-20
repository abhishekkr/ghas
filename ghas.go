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
	sizeToHash := flag.Int("s", 64, "size of hash to generate")
	fileToHash := flag.String("f", "", "path of file to hash")
	flag.Parse()
	if *fileToHash != "" {
		hashTheFile(*fileToHash, *sizeToHash)
	}
}

func checkFatal(e error) {
	if e != nil {
		panic(e)
	}
}

func hashTheFile(filepath string, hashSize int) {
	dat, err := ioutil.ReadFile(filepath)
	fmt.Println("file length: ", len(dat))
	checkFatal(err)
	fmt.Println("\n-GHAS->")
	start := time.Now()
	ghash := GetGhas(dat, hashSize)
	fmt.Println(time.Since(start))
	fmt.Println(ghash)
	fmt.Println(len(ghash))
	otherHashing(dat)
}

func GetGhas(dat []byte, hashSize int) string {
	g := ghaslib.New(hashSize)
	g.Eval([]byte(dat))
	return g.String()
}

func otherHashing(dat []byte) {
	fmt.Println("\n-MD5->")
	mstart := time.Now()
	mhashB := md5.Sum(dat)
	mhash := hex.EncodeToString(mhashB[:])
	fmt.Println(time.Since(mstart))
	fmt.Println(mhash)
	fmt.Println(len(mhash))

	fmt.Println("\n-SHA256->")
	s2start := time.Now()
	s256B := sha256.Sum256(dat)
	s256 := hex.EncodeToString(s256B[:])
	fmt.Println(time.Since(s2start))
	fmt.Println(s256)
	fmt.Println(len(s256))

	fmt.Println("\n-SHA512->")
	s5start := time.Now()
	s512B := sha512.Sum512(dat)
	s512 := hex.EncodeToString(s512B[:])
	fmt.Println(time.Since(s5start))
	fmt.Println(s512)
	fmt.Println(len(s512))

	fmt.Println("\n-HMAC512->")
	hstart := time.Now()
	hmac512 := hmac.New(sha512.New, []byte("secret"))
	hmac512.Write(dat)
	fmt.Printf("hmac512:\t%s\n", base64.StdEncoding.EncodeToString(hmac512.Sum(nil)))
	fmt.Println(time.Since(hstart))
}
