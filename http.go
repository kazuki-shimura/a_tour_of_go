package main

import (
	"fmt"
	"net/http"
)

// 簡単なWeb上での表示
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello HTTPサーバー")
}
func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<div><h1>田中</h1></div>")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/2", handler2)
	http.ListenAndServe(":8080", nil)
}

