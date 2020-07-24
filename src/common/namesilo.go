package common

import (
	"encoding/xml"
	"log"
	"net/http"
	"net/url"
)

type Response struct {
	Request Request `xml:"request"`
	Reply   Reply   `xml:"reply"`
}

type Request struct {
	Operation string `xml:"operation"`
	Ip        string `xml:"ip"`
}

type Reply struct {
	Code               int              `xml:"code"`
	Detail             string           `xml:"detail"`
	ResourceRecordList []ResourceRecord `xml:"resource_record"`
}

type ResourceRecord struct {
	RecordId string `xml:"record_id"`
	Type     string `xml:"type"`
	Host     string `xml:"host"`
	Value    string `xml:"value"`
	Ttl      string `xml:"ttl"`
	Distance string `xml:"distance"`
}

var key string
var domain0 string

func HandleDns(targetHost string, apiKey string, domain string) {
	key = apiKey
	domain0 = domain
	update, id, localIp := queryDnsRecordIdAndLocalIp(targetHost + "." + domain0)
	if update {
		if id == "" {
			log.Println("没查到id")
			addDnsRecord(localIp, targetHost)
		} else {
			updateDnsRecord(id, localIp, targetHost)
		}
	}
}

func queryDnsRecordIdAndLocalIp(targetHost string) (bool, string, string) {
	host := targetHost
	query := url.Values{}
	log.Printf("查询DNS记录%s", host)
	response := requestDnsServer("dnsListRecords", query)
	id, dnsIp := findRecordIdAndIpByHost(host, response.Reply.ResourceRecordList)
	log.Printf("DNS记录%s的id为:%s", host, id)
	if dnsIp == response.Request.Ip {
		log.Println("dnsIp和当前ip一致,不需要更新dns记录")
		return false, "", ""
	} else {
		return true, id, response.Request.Ip
	}
}

func updateDnsRecord(recordId string, ip string, targetHost string) {
	query := url.Values{}
	query.Add("rrid", recordId)
	query.Add("rrvalue", ip)
	query.Add("rrhost", targetHost)
	query.Add("rrttl", "3603")
	log.Printf("修改DNS记录%s的ip为%s", recordId, ip)
	response := requestDnsServer("dnsUpdateRecord", query)
	log.Println(response)
}

func addDnsRecord(ip string, targetHost string) {
	query := url.Values{}
	query.Add("rrvalue", ip)
	query.Add("rrhost", targetHost)
	query.Add("rrttl", "3603")
	query.Add("rrtype", "A")
	log.Printf("新增DNS记录ip为%s", ip)
	response := requestDnsServer("dnsAddRecord", query)
	log.Println(response)
}

func requestDnsServer(operation string, params url.Values) Response {
	request := http.Request{}
	requestUrl := url.URL{}
	request.URL = &requestUrl
	requestUrl.Scheme = "http"
	requestUrl.Host = "www.namesilo.com"
	requestUrl.Path = "/api/" + operation
	params.Add("version", "1")
	params.Add("type", "xml")
	params.Add("key", key)
	params.Add("domain", domain0)
	requestUrl.RawQuery = params.Encode()
	data, _ := Get(request)
	return parseResponse(data)
}

func parseResponse(body []byte) Response {
	response := Response{}
	err := xml.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}
	if response.Reply.Code != 300 {
		log.Fatal(response.Reply.Detail)
	}
	return response
}

func findRecordIdAndIpByHost(host string, resourceRecordList []ResourceRecord) (string, string) {
	if len(resourceRecordList) == 0 {
		return "", ""
	}
	for _, resourceRecord := range resourceRecordList {
		if resourceRecord.Host == host {
			return resourceRecord.RecordId, resourceRecord.Value
		}
	}
	return "", ""
}
