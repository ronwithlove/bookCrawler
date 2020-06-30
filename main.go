package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.99comic.com")
	if err!=nil{
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode!=http.StatusOK{
		fmt.Printf("Error status code:%d",resp.StatusCode)
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		panic(err)
	}

	fmt.Printf("%s",result)
}
