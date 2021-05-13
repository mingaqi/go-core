package main

import (
	"net/http"
)

func main() {

	master := http.NewServeMux()
	// 代理处理
	master.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("success"))
	})
	http.ListenAndServe(":8081", master)

}
