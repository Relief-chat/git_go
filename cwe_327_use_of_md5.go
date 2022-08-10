package main

import (
	"crypto/des"
	"crypto/md5"
	"crypto/sha3"
	"crypto/rc4"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
}



func test_md5() {
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	defer func() {
		err := f.Close()
		if err != nil {
			log.Printf("error closing the file: %s", err)
		}
	}()

	// bad
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	// bad
	fmt.Printf("%x", md5.Sum(nil))
}


func test_sha3() {
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	defer func() {
		err := f.Close()
		if err != nil {
			log.Printf("error closing the file: %s", err)
		}
	}()

	// ok
	h := sha3.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	// ok
	fmt.Printf("%x", sha3.Sum(nil))
}

