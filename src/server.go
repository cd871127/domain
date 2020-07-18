package main

import (
	"common"
	"log"
	"net/http"
	"os"
	"strings"
)

type ServerConfig struct {
	port     string
	logFile  string
	password string
	prefix   string
}

var ipAddr = "UNKNOWN"
var serverConfig ServerConfig
//nohup sudo -u app /app/domain_server /app/config/  > /dev/null 2>&1 &
func main() {
	var configPath string
	if len(os.Args) >= 2 {
		configPath = os.Args[1]
	} else {
		configPath = "config/"
	}
	log.Println("init server...")
	serverConfig = loadConfig(configPath)
	common.InitLogger(serverConfig.logFile)

	http.HandleFunc("/", index)
	http.HandleFunc("/ip", ip)
	log.Println("start server...")
	log.Fatal(http.ListenAndServe(":"+serverConfig.port, nil))
}

func ip(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	remoteAddr := r.RemoteAddr
	if !strings.Contains(remoteAddr, serverConfig.prefix) {
		log.Println("未知ip：" + remoteAddr)
	}
	password := r.URL.Query().Get("password")
	if password == serverConfig.password {
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

func loadConfig(configPath string) ServerConfig {
	var serverConfig = ServerConfig{}
	v, _ := common.Load(configPath)
	serverConfig.logFile = v.GetString("server.logFile")
	serverConfig.prefix = v.GetString("server.prefix")
	serverConfig.password = v.GetString("server.password")
	serverConfig.port = v.GetString("server.port")
	log.Println(serverConfig)
	return serverConfig
}
