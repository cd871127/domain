package main

import (
	"common"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	param := url.Values{}
	param.Add("test", "123")
	param.Add("test2", "123")
	common.Get("ip.shit-code.com", "123/123", "50888", param)
}

//注册ip
func registerIp(host string, port int, ip string) {
	passwd := "cd123321"
	client := &http.Client{}

	fmt.Println("http://" + host + ":" + strconv.Itoa(port) + "/ip?ip=" + ip + "&passwd=" + passwd)
	//resp, _ := client.Get( host + ":" + string(port) + "/ip?ip=" + ip + "&passwd=" + passwd)
	resp, _ := client.Get("http://localhost:8888/ip?ip=ffffaf&passwd=1")
	if resp != nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(body))
	}

}

func get1() {
	client := &http.Client{}
	resp, _ := client.Get("http://localhost:8888")
	if resp != nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(body))
	}
}
