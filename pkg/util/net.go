package util

import (
	"io/ioutil"
	"net/http"
	"time"
)

/*
带有lingjiang头信息的Get请求
url ：传入url信息
返回
body:接收的数据
err:错误信息
*/
func Get(url string) (body []byte, err error) {
	client := &http.Client{Timeout: time.Second * 10}
	var rqt *http.Request
	rqt, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	//灵匠标志
	rqt.Header.Add("User-Agent", "Lingjiang")
	var response *http.Response
	response, err = client.Do(rqt)
	if err != nil {
		return
	}
	defer func() {
		response.Body.Close()
	}()
	body, err = ioutil.ReadAll(response.Body)
	return
}
