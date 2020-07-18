package common

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var client = &http.Client{}

func Get(host string, path string, port string, param url.Values) string {
	serverUrl := url.URL{}
	serverUrl.Path = path
	serverUrl.Host = host + ":" + port
	serverUrl.Scheme = "http"
	if param != nil {
		serverUrl.RawQuery = param.Encode()
	}
	log.Printf("%s", serverUrl.String())
	request, err := http.NewRequest("GET", serverUrl.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
	}
	if response != nil {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		return string(body)
	}
	return ""
}
