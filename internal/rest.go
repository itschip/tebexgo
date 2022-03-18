package internal

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func GetRequest(secret string, endpoint string) ([]byte, error) {
	reqURL, _ := url.Parse(endpoint)

	req := &http.Request{
		Method: "GET",
		URL:    reqURL,
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

// Should change this to use DefaultClient
func PostRequest(secret string, endpoint string, body []byte) ([]byte, error) {
	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		log.Println("Failed to post request")
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Tebex-Secret", secret)

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	fmt.Println(response.Status)

	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return body, nil
}

func PutRequest(secret string, endpoint string, body []byte) ([]byte, error) {
	request, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(body))
	if err != nil {
		log.Println("Failed to post request")
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Tebex-Secret", secret)

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	fmt.Println(response.Status)

	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return body, nil
}

func DeleteRequest(secret string, endpoint string, body []byte) ([]byte, error) {
	request, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		log.Println("Failed to post request")
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Tebex-Secret", secret)

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	fmt.Println(response.Status)

	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return body, nil
}
