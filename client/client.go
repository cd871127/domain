package main

import (
	"domain/common"
	"log"
	"net/http"
	"net/url"
	"os"
)

type config struct {
	ClientConfig ClientConfig `mapstructure:"client"`
}

type ClientConfig struct {
	LogFile  string
	Namesilo Namesilo
	Server   Server
}

type Namesilo struct {
	ApiKey     string `mapstructure:"api-key"`
	TargetHost string
	Domain     string
}

type Server struct {
	Port     string
	Host     string
	Password string
}

var clientConfig ClientConfig

func main() {
	var configPath string
	if len(os.Args) >= 2 {
		configPath = os.Args[1]
	} else {
		configPath = "config/"
	}
	log.Printf("log path:%s", configPath)
	log.Println("init client...")

	clientConfig = loadClientConfig(configPath)
	common.InitLogger(clientConfig.LogFile)
	localIp := common.HandleDns(clientConfig.Namesilo.TargetHost, clientConfig.Namesilo.ApiKey, clientConfig.Namesilo.Domain)
	registerIp(clientConfig.Server, localIp)
}

func loadClientConfig(configPath string) ClientConfig {
	//var clientConfig = ClientConfig{}
	//clientConfig.server = Server{}
	//clientConfig.namesilo = Namesilo{}

	//v, _ := common.Load(configPath)


	var config = config{}
	v, _ := common.Load(configPath)
	v.Unmarshal(&config)
	log.Println(config)
	return config.ClientConfig

	//clientConfig.logFile = v.GetString("client.logFile")
	//clientConfig.server.host = v.GetString("client.server.host")
	//clientConfig.server.password = v.GetString("client.server.password")
	//clientConfig.server.port = v.GetString("client.server.port")
	//clientConfig.namesilo.apiKey = v.GetString("client.namesilo.api-key")
	//clientConfig.namesilo.apiKey = v.GetString("client.namesilo.api-key")
	//clientConfig.namesilo.targetHost = v.GetString("client.namesilo.targetHost")
	//clientConfig.namesilo.domain = v.GetString("client.namesilo.domain")
	//log.Println(clientConfig)
	//return clientConfig
}

//注册ip
func registerIp(server Server, localIp string) {
	request := http.Request{}
	requestUrl := url.URL{}
	request.URL = &requestUrl
	requestUrl.Scheme = "http"
	requestUrl.Host = server.Host + ":" + server.Port
	requestUrl.Path = "/ip"
	params := url.Values{}
	//param.Add("", "123")
	params.Add("password", server.Password)
	params.Add("ip", localIp)
	requestUrl.RawQuery = params.Encode()
	body, _ := common.Get(request)
	log.Printf("注册IP结果：%s", string(body))
}
