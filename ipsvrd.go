//做成查询服务
//在gopath的src目录下创建golang.org/x目录，进入到golang.org/x目录，执行命令：
//git clone https://github.com/golang/net.git
//git clone https://github.com/golang/text.git
package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"encoding/json"
)

var (
	port     = 8081
	qstr = "http://whois.pconline.com.cn/ipJson.jsp?ip=%s&json=true"
)

type IpInfo struct{
	Ip string `json:"ip"`
	Pro   string `json:"pro"`
	ProCode string `json:"proCode"`
	City string `json:"city"`
	CityCode string `json:"cityCode"`
	Region string `json:"region"`
	RegionCode string `json:"regionCode"`
	Addr string `json:"addr"`
	RegionNames string `regionNames:"addr"`
	Err string `regionNames:"err"`
}

func IpJsontoStruct(ipstring string) string{
	p := &IpInfo{}
	err := json.Unmarshal([]byte(ipstring), p)
	if err != nil {
		fmt.Println("ipJsontoStruct error",ipstring)
		return p.Addr
	}
	return p.Addr
}

func QueryIpInfo(ip string) string {
	querystr := fmt.Sprintf(qstr,ip)
	resp, err := http.Get(querystr)
    if err != nil {
        fmt.Println(err)
        return ""
    }
    defer resp.Body.Close()
    utf8Reader := transform.NewReader(resp.Body, 
		simplifiedchinese.GBK.NewDecoder())
	body, err := ioutil.ReadAll(utf8Reader)
    addr := IpJsontoStruct(string(body))
    return addr
}

func handler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	if r.Method != "GET" {
		return
	}
	if r.URL.Path == "/favicon.ico" {
		return
	}
	ip := r.URL.Path[1:]
	fmt.Println("ip",ip)
	ipAddr := QueryIpInfo(ip)
	fmt.Println(ipAddr)
	w.Write([]byte(ipAddr))
}

func main(){
	server := http.Server{
		Addr:           fmt.Sprintf(":%d", port),
	}
	http.HandleFunc("/",handler)
	server.ListenAndServe()
}