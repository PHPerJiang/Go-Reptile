package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp,err := http.Get("http://www.zhenai.com/zhenghun");
	if err != nil {
		panic("未抓取到网站");
	}
	//关闭连接
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code",resp.StatusCode)
		return
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n",all)
}
