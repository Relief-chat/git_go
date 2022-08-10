package main

import (
    "crypto/md5"
    "crypto/sha256"
    "fmt"
    "io"
)

//// True positives ////
func ex1(user *User, pwtext string) {
    h := md5.New()
    io.WriteString(h, pwtext)
    // bad
    user.setPassword(h.Sum(nil))
}

func ex2(user *User, pwtext string) {
    data := []byte(pwtext)
    // bad
    user.setPassword(md5.Sum(data))
}

//// True negatives ////
func ok1(user *User, pwtext string) {
    h := sha256.New()
    io.WriteString(h, pwtext)
    // ok 
    user.setPassword(h.Sum(nil))
}

func ok2(user *User, pwtext string) {
    data := []byte(pwtext)
    // ok 
    user.setPassword(sha256.Sum(data))
}

func ok3(user *User, pwtext string) {
    data := []byte(pwtext)
    // ok 
    user.setSomethingElse(md5.Sum(data))
}
