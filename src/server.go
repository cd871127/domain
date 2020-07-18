package main

import (
	"config"
	"log"
	"net/http"
	"strings"
)

var ipAddr = "UNKNOWN"

func main() {
	serverConfig, _ := config.Load()
	port := serverConfig.GetString("server.port")
	http.HandleFunc("/", index)
	http.HandleFunc("/ip", ip)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func ip(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	remoteAddr := r.RemoteAddr
	serverConfig, _ := config.Load()
	configPassword := serverConfig.GetString("server.password")
	prefix := serverConfig.GetString("server.prefix")
	if !strings.Contains(remoteAddr, prefix) {
		log.Println("未知ip：" + remoteAddr)
	}
	password := r.URL.Query().Get("password")
	if password == configPassword {
		ipAddr = r.URL.Query().Get("ip")
		if ipAddr == "" {
			ipAddr = remoteAddr
		}
		_, _ = w.Write([]byte("ok"))
		return
	}
	_, _ = w.Write([]byte("fail"))
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	_, _ = w.Write([]byte(ipAddr))
}
