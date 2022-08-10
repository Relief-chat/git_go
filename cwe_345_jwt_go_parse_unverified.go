package main

import (
    "fmt"

    "github.com/dgrijalva/jwt-go"
)

func bad1(tokenString string) {
    // bad
    token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
    if err != nil {
        fmt.Println(err)
        return
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok {
        fmt.Println(claims["foo"], claims["exp"])
    } else {
        fmt.Println(err)
    }
}

func ok1(tokenString string, keyFunc Keyfunc) {
    // ok 
    token, err := new(jwt.Parser).ParseWithClaims(tokenString, jwt.MapClaims{}, keyFunc)
    if err != nil {
        fmt.Println(err)
        return
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok {
        fmt.Println(claims["foo"], claims["exp"])
    } else {
        fmt.Println(err)
    }
}
