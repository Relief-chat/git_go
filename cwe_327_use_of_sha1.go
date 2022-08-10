package main

import (
	"crypto/des"
	"crypto/md5"
	"crypto/rc4"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
}

func test_sha1() {
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// bad
	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	// bad
	fmt.Printf("%x", sha1.Sum(nil))
}

func test_sha256() {
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// ok
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	// ok
	fmt.Printf("%x", sha256.Sum(nil))
}
