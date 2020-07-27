package main

import (
	"domain/common"
	"log"
	"net/http"
	"os"
	"strings"
)

type ServerConfig struct {
	Port     string
	LogFile  string
	Password string
	Prefix   string
}
type config struct {
	Server ServerConfig `mapstructure:"server"`
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
	serverConfig = loadServerConfig(configPath)
	common.InitLogger(serverConfig.LogFile)

	http.HandleFunc("/", index)
	http.HandleFunc("/ip", ip)
	log.Println("start server...")
	log.Fatal(http.ListenAndServe(":"+serverConfig.Port, nil))
}

func ip(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	remoteAddr := r.RemoteAddr
	if !strings.Contains(remoteAddr, serverConfig.Prefix) {
		log.Println("未知ip：" + remoteAddr)
	}
	password := r.URL.Query().Get("password")
	if password == serverConfig.Password {
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

func loadServerConfig(configPath string) ServerConfig {
	var config = config{}
	v, _ := common.Load(configPath)
	v.Unmarshal(&config)
	log.Println(config)
	return config.Server
}
