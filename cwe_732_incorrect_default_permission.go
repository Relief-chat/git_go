package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
}

func test_chmod() {
	// bad
	err := os.Chmod("/tmp/somefile", 0777)
	if err != nil {
		fmt.Println("Error when changing file permissions!")
		return
	}

	// ok 
	err := os.Chmod("/tmp/somefile", 0400)
	if err != nil {
		fmt.Println("Error when changing file permissions!")
		return
	}
}

func test_mkdir() {
	// bad
	err := os.Mkdir("/tmp/mydir", 0777)
	if err != nil {
		fmt.Println("Error when creating a directory!")
		return
	}

	// bad
	err := os.MkdirAll("/tmp/mydir", 0777)
	if err != nil {
		fmt.Println("Error when creating a directory!")
		return
	}

	// ok 
	err := os.MkdirAll("/tmp/mydir", 0600)
	if err != nil {
		fmt.Println("Error when creating a directory!")
		return
	}
}

func test_openfile() {
	// bad
	_, err := os.OpenFile("/tmp/thing", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error opening a file!")
		return
	}

	// ok 
	_, err := os.OpenFile("/tmp/thing", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("Error opening a file!")
		return
	}
}

func test_writefile() {
	// bad
	err := ioutil.WriteFile("/tmp/demo2", []byte("This is some data"), 0644)
	if err != nil {
		fmt.Println("Error while writing!")
	}
}
