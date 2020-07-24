package main

import (
	"common"
	"log"
	"net/http"
	"net/url"
	"os"
)

type ClientConfig struct {
	logFile  string
	namesilo Namesilo
	server   Server
}

type Namesilo struct {
	apiKey     string
	targetHost string
	domain     string
}

type Server struct {
	port     string
	host     string
	password string
}

var clientConfig ClientConfig

func main() {
	var configPath string
	if len(os.Args) >= 2 {
		configPath = os.Args[1]
	} else {
		configPath = "config/"
	}
	log.Println("init client...")

	clientConfig = loadClientConfig(configPath)
	common.InitLogger(clientConfig.logFile)
	localIp := common.HandleDns(clientConfig.namesilo.targetHost, clientConfig.namesilo.apiKey, clientConfig.namesilo.domain)
	registerIp(clientConfig.server, localIp)
}

func loadClientConfig(configPath string) ClientConfig {
	var clientConfig = ClientConfig{}
	clientConfig.server = Server{}
	clientConfig.namesilo = Namesilo{}

	v, _ := common.Load(configPath)

	clientConfig.logFile = v.GetString("client.logFile")
	clientConfig.server.host = v.GetString("client.server.host")
	clientConfig.server.password = v.GetString("client.server.password")
	clientConfig.server.port = v.GetString("client.server.port")
	clientConfig.namesilo.apiKey = v.GetString("client.namesilo.api-key")
	clientConfig.namesilo.apiKey = v.GetString("client.namesilo.api-key")
	clientConfig.namesilo.targetHost = v.GetString("client.namesilo.targetHost")
	clientConfig.namesilo.domain = v.GetString("client.namesilo.domain")
	log.Println(clientConfig)
	return clientConfig
}

//注册ip
func registerIp(server Server, localIp string) {
	request := http.Request{}
	requestUrl := url.URL{}
	request.URL = &requestUrl
	requestUrl.Scheme = "http"
	requestUrl.Host = server.host + ":" + server.port
	requestUrl.Path = "/ip"
	params := url.Values{}
	//param.Add("", "123")
	params.Add("password", server.password)
	params.Add("ip", localIp)
	requestUrl.RawQuery = params.Encode()
	body, _ := common.Get(request)
	log.Printf("注册IP结果：%s", string(body))
}
