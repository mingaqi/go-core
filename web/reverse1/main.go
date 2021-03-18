package main

import (
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	http.HandleFunc("/", ServeHTTP)
	http.ListenAndServe(":8086", nil)
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/favicon.ico" {
		io.WriteString(w, "Request path Error")
		return
	}
	remote, err := url.Parse("http://" + "127.0.0.1:8081")
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}
