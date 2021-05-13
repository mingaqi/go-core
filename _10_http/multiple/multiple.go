package main

import (
	"net/http"
	"time"
)

func main() {
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/mux1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello mux1"))
	})
	go http.ListenAndServe(":8081", mux1)

	//----------------------------------

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/mux2", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello mux2"))
	})
	server := http.Server{
		Addr:         ":8082",
		Handler:      mux2,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go server.ListenAndServe()

	select {}
}
