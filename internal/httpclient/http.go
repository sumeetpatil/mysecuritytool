package httpclient

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

type HttpClient struct {
	headers map[string]string
	url     string
}

type HttpClientHandler interface {
	Get()
	Post(body string)
}

type HttpResp struct {
	StatusCode int
	Body       string
}

func NewHttpClient(url string, headers map[string]string) HttpClient {
	return HttpClient{url: url, headers: headers}
}

func (httpClient *HttpClient) Get() HttpResp {
	return call(httpClient, "")
}

func (httpClient *HttpClient) Post(body string) HttpResp {
	return call(httpClient, body)
}

func call(httpClient *HttpClient, body string) HttpResp {
	req, err := http.NewRequest("GET", httpClient.url, nil)
	if body != "" {
		req, err = http.NewRequest("POST", httpClient.url, strings.NewReader(body))
	}

	if err != nil {
		fmt.Println("Error creating request:", err)
		return HttpResp{StatusCode: 500}
	}

	for k, v := range httpClient.headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return HttpResp{StatusCode: 500}
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Request failed. Status code:", response.StatusCode)
		return HttpResp{StatusCode: 500}
	}

	var responseBody bytes.Buffer
	_, err = responseBody.ReadFrom(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return HttpResp{StatusCode: 500}
	}

	return HttpResp{StatusCode: 200, Body: responseBody.String()}
}
