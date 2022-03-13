package internal

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetRequest(secret string, endpoint string) ([]byte, error) {
	reqURL, _ := url.Parse(endpoint)
	
	req := &http.Request{
		Method: "GET",
		URL: reqURL,
		Header: map[string][]string{
			"X-Tebex-Secret": {secret},
		},
	}
	
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	
	return response, nil
}

func PostRequest(secret string, endpoint string, body io.ReadCloser) ([]byte, error) {
	reqURL, _ := url.Parse(endpoint)
	
	req := &http.Request{
		Method: "GET",
		URL: reqURL,
		Body: body,
		Header: map[string][]string{
			"X-Tebex-Secret": {secret},
		},
	}
	
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	
	return response, nil
}