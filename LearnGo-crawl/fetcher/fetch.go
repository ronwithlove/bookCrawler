package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//提取网页内容
func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code:%d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

//模拟游览器访问
func BrowerFetch(url string) ([]byte, error) {
	client:=&http.Client{}
	req,err:=http.NewRequest("Get",url,nil)
	if err!=nil{
		return nil,fmt.Errorf("ERROR:get url:%s",url)
	}

	req.Header.Set("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body,err:=ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Printf("%s",body)

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code:%d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}