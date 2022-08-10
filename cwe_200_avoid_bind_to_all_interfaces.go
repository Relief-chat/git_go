package main

import (
	"log"
	"net"
)

func bind_all() {
	// bad
	l, err := net.Listen("tcp", "0.0.0.0:2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
}

func bind_default() {
	// bad
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
}

func main() {
	// ok 
	l, err := net.Listen("tcp", "192.168.1.101:2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
}
