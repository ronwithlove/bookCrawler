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

