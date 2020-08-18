package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/index", index)
	if err := http.ListenAndServe("localhost:9090", nil); err != nil {
		log.Fatal("http服务启动失败")
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	str := `<h1 style="color:red">hello, net.http<h1>`
	w.Write([]byte(str))
}
