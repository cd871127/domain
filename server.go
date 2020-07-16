package main

import (
	"log"
	"net/http"
)

var ipAddr = ""

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/ip", ip)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func ip(w http.ResponseWriter, r *http.Request) {
	passwd := r.URL.Query().Get("passwd")
	if passwd == "1" {
		ipAddr = r.URL.Query().Get("ip")
		if ipAddr == "" {
			ipAddr = r.RemoteAddr
		}
		_, _ = w.Write([]byte("ok"))
		return
	}
	_, _ = w.Write([]byte("fail"))
}

func index(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(ipAddr))
}
