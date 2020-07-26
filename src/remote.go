package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"syscall"
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
	command := exec.Command("/usr/bin/wakeonlan -i 192.168.2.0 d0:50:99:70:87:3f")
	outinfo := bytes.Buffer{}
	command.Stdout = &outinfo
	err := command.Start()
	if err != nil {
		fmt.Println(err.Error())
	}
	if err = command.Wait(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(command.ProcessState.Pid())
		fmt.Println(command.ProcessState.Sys().(syscall.WaitStatus).ExitStatus())
		fmt.Println(outinfo.String())
	}
}
