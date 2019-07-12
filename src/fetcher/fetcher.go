package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string)([]byte, error) {
	//由于get会遇到403所以换方式
	//resp,err := http.Get(url);

	client := &http.Client{}
	req, err := http.NewRequest("GET",url,nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil{
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	//出错返回错误码
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code",resp.StatusCode)
		return nil,fmt.Errorf("wrong status code: %d",resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}
