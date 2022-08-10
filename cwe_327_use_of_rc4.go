package main

import (
	"crypto/des"
	"crypto/md5"
	"crypto/rc4"
	"crypto/aes"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
}

func test_rc4() {
	key := []byte{1, 2, 3, 4, 5, 6, 7}
	// bad
	c, err := rc4.NewCipher(key)
	dst := make([]byte, len(src))
	c.XORKeyStream(dst, src)
}

func test_aes() {
	key := []byte{1, 2, 3, 4, 5, 6, 7}
	// ok
	c, err := aes.NewCipher(key)
	dst := make([]byte, len(src))
	c.XORKeyStream(dst, src)
}
