package main

import (
	"crypto/des"
	"crypto/md5"
	"crypto/rc4"
	"crypto/sha1"
	"crypto/aes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
}

func test_des() {
	ede2Key := []byte("example key 1234")

	var tripleDESKey []byte
	tripleDESKey = append(tripleDESKey, ede2Key[:16]...)
	tripleDESKey = append(tripleDESKey, ede2Key[:8]...)
	// bad
	_, err := des.NewTripleDESCipher(tripleDESKey)
	if err != nil {
		panic(err)
	}
}

func test_des_ok() {
	ede2Key := []byte("example key 1234")

	var tripleDESKey []byte
	tripleDESKey = append(tripleDESKey, ede2Key[:16]...)
	tripleDESKey = append(tripleDESKey, ede2Key[:8]...)
	// bad
	_, err := aes.NewTripleDESCipher(tripleDESKey)
	if err != nil {
		panic(err)
	}
}
