package uhoh

import (
	"context"
	"net"
	"net/http"
	"net/http/httptrace"
)

func WithTrace(req *http.Request, trace *httptrace.ClientTrace) *http.Request {
    // bad
	return req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
}
