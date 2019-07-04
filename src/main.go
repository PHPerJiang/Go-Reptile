package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
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
	//fmt.Printf("%s\n",all)
	getCityList(all);
}

/**
获取城市列表
 */
func getCityList(contents []byte)  {
	//匹配链接
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" data-v-5e16505f>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents,-1)
	for _, m := range matches{
		fmt.Printf("City: %s, Url: %s\n", m[2],m[1])
	}
	fmt.Printf("has find :%d\n",len(matches))
}


