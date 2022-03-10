package httplus

import (
	"net"
	"net/http"
	"strconv"
)

// GetHostPort returns host and port from a proxy HTTP request
func GetHostPort(r *http.Request) (host string, port int) {
	host = r.Host
	host, sport, err := net.SplitHostPort(r.Host)
	if err == nil {
		port, _ = strconv.Atoi(sport)
	} else {
		if r.Host != "" {
			host = r.Host
		}
		if r.URL.Scheme == "" || r.URL.Scheme == "http" {
			port = 80
		}
		if r.URL.Scheme == "https" {
			port = 443
		}
	}
	return host, port
}

// CopyHeader copies all the headers from src to dst
func CopyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
