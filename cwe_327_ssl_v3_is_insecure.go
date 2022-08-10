package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)
 
type zeroSource struct{}

func (zeroSource) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = 0
	}

	return len(b), nil
}

func main() {
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	server.TLS = &tls.Config{
		Rand: zeroSource{}, 
	}
	server.StartTLS()
	defer server.Close()
 
	w := os.Stdout

	client := &http.Client{
		Transport: &http.Transport{
			// bad
			TLSClientConfig: &tls.Config{
				KeyLogWriter:       w,
				MinVersion:         tls.VersionSSL30,
				Rand:               zeroSource{},  
				InsecureSkipVerify: true,          
			},
		},
	}
	resp, err := client.Get(server.URL)
	if err != nil {
		log.Fatalf("Failed to get URL: %v", err)
	}
	resp.Body.Close()

	client_good := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				KeyLogWriter: w,
				// OK
				MinVersion:         tls.VersionTLS10,
				Rand:               zeroSource{},  
				InsecureSkipVerify: true,          
			},
		},
	}
	resp, err := client.Get(server.URL)
	if err != nil {
		log.Fatalf("Failed to get URL: %v", err)
	}
	resp.Body.Close()
}
