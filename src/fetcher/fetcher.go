package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string)([]byte, error) {
	resp,err := http.Get(url);
	if err != nil {
		return nil, err
	}
	//关闭连接
	defer resp.Body.Close()

	//出错返回错误码
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code",resp.StatusCode)
		return nil,fmt.Errorf("wrong status code: %d",resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}
