package insecuregrpc

import (
	"crypto/x509"
	"log"
	"net/http"
	"net/http/httptest"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// cf. https://blog.gopheracademy.com/advent-2019/go-grps-and-tls/#connection-without-encryption
func unsafe() {
	// bad
	s := grpc.NewServer()
 
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func safe() {
 
	// ok 
	s := grpc.NewServer(grpc.Creds(credentials.NewClientTLSFromCert(x509.NewCertPool(), "")))
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
