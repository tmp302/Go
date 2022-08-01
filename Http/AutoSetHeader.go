package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func AutoSetHeaders(Headers string) ([]string,[]string){
	var (
		keyResult []string
		valResult []string
	)

	requestInfo := Headers
	allInfo:= strings.Split(requestInfo, "\n\n")
	headers := strings.Split(allInfo[0], "\n")

	for _,i:= range headers[1:]{
		i = strings.Replace(i, " ", "", -1)
		key := strings.SplitN(i, ":", 2)
		keyResult = append(keyResult, key[0])
		valResult = append(valResult, key[1])
	}
	return keyResult,valResult
}

func main(){
	key,val := AutoSetHeaders(`GET /test HTTP/1.1
    Host: xxx.xxx.com
    Accept: application/json, text/plain, */*
    Sec-Ch-Ua-Mobile: ?0
    User-Agent: xxxxxx
    Origin: https://xxx.xxx.com
    Sec-Fetch-Site: same-site
    Sec-Fetch-Mode: cors
    Sec-Fetch-Dest: empty
    Referer: https://xxx.xxx.com/
    Accept-Encoding: gzip, deflate
    Accept-Language: zh-CN,zh;q=0.9
    Connection: close`)

	request, err := http.NewRequest(http.MethodGet, "https://httpbin.org/get", nil)
	if err != nil{
		panic(err)
	}
	for count,i := range key{
		request.Header.Add(i, val[count])
	}
	// 执行
	r, err := http.DefaultClient.Do(request)
	if err != nil{
		panic(err)
	}
	defer func() {_ = r.Body.Close()}()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil{
		panic(err)
	}
	fmt.Println((string)(content))
}