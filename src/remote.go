package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	//var configPath string
	//if len(os.Args) >= 2 {
	//	configPath = os.Args[1]
	//} else {
	//	configPath = "config/"
	//}
	//log.Println("init server...")
	//serverConfig = loadServerConfig(configPath)
	//common.InitLogger(serverConfig.logFile)

	http.HandleFunc("/", start)
	log.Println("start server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func start(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("wakeonlan", "-i", "192.168.2.0 d0:50:99:70:87:3f")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Execute Command failed:" + err.Error())
		return
	}
	fmt.Println("Execute Command finished.")
}
