package main

import (
	"crypto/rand"
	mrand "math/rand"
)

func main() {
	main0()
	main1()
	main2()
	main3()
}

func main0() {
	// bad
	bad, _ := mrand.Read(nil)
	println(bad)
}

func main1() {
	// ok 
	good, _ := rand.Read(nil)
	println(good)
}

func main2() {
	// bad
	bad := mrand.Int()
	println(bad)
}

func main3() {
	// ok 
	good, _ := rand.Read(nil)
	println(good)
	i := mrand.Int31()
	println(i)
}
