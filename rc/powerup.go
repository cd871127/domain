package main

import (
	"bytes"
	"domain/common"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"syscall"
)

type RemoteControlConfig struct {
	Port    string
	LogFile string
	Client  struct {
		Ip  string
		Mac string
	}
}
type config struct {
	RemoteControl RemoteControlConfig `mapstructure:"remote"`
}

var remoteControlConfig RemoteControlConfig

func main() {
	var configPath string
	if len(os.Args) >= 2 {
		configPath = os.Args[1]
	} else {
		configPath = "config/"
	}
	log.Println("init server...")
	remoteControlConfig = loadRemoteConfig(configPath)
	common.InitLogger(remoteControlConfig.LogFile)

	http.HandleFunc("/", start)
	log.Println("start server...")
	log.Fatal(http.ListenAndServe(":"+remoteControlConfig.Port, nil))
}

func start(w http.ResponseWriter, r *http.Request) {
	cmd := "wakeonlan -i " + remoteControlConfig.Client.Ip + " " + remoteControlConfig.Client.Mac

	command := exec.Command("/bin/sh", "-c", cmd)
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

func loadRemoteConfig(configPath string) RemoteControlConfig {
	var config = config{}
	v, _ := common.Load(configPath)
	v.Unmarshal(&config)
	log.Println(config)
	return config.RemoteControl
}
