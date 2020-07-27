package common

import (
	"io/ioutil"
	"net/http"
)

func Get(request http.Request) ([]byte, error) {
	//查询dns列表
	client := &http.Client{}
	resp, err := client.Do(&request)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
