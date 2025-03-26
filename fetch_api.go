package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var url string = "https://valorant-api.com/v1/agents/"
	var method string = "GET"
	var headers map[string]string = map[string]string{
		"Accept": "application/json",
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println(string(body))
}
