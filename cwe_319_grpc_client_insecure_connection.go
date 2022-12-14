package insecuregrpc

import (
    "google.golang.org/grpc"
)

// cf. https://blog.gopheracademy.com/advent-2019/go-grps-and-tls/#connection-without-encryption
func unsafe() {
    // bad
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
}

func safe() {
    // ok 
    conn, err := grpc.Dial(address)
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
}
