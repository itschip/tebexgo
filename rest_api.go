package tebexgo

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (s *Session) GetRequest(endpoint string) ([]byte, error) {
	reqURL, _ := url.Parse(endpoint)
	
	req := &http.Request{
		Method: "GET",
		URL: reqURL,
		Header: map[string][]string{
			"X-Tebex-Secret": {s.Secret},
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

func (s *Session) PostRequest(endpoint string, body io.ReadCloser) ([]byte, error) {
	reqURL, _ := url.Parse(endpoint)
	
	req := &http.Request{
		Method: "GET",
		URL: reqURL,
		Body: body,
		Header: map[string][]string{
			"X-Tebex-Secret": {s.Secret},
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