package getui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var timeout time.Duration

func SetTimeout(t time.Duration) {
	timeout = t
}

func newClient() *http.Client {
	if timeout.Nanoseconds() <= 0 {
		timeout = defaultRequestTimeout
	}
	return &http.Client{
		Timeout: timeout,
	}
}

func newRequest(token string, method string, url string, body interface{}) (*http.Request, error) {
	headers := http.Header{
		"Content-Type": []string{"application/json;charset=utf-8"},
	}
	if len(token) > 0 {
		headers.Add("token", token)
	}

	content, err := json.Marshal(body)
	if err != nil {
		return &http.Request{}, fmt.Errorf("json.Marshal fail")
	}

	var request *http.Request

	if body != nil {
		request, err = http.NewRequest(method, url, bytes.NewBuffer(content))
	} else {
		request, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return &http.Request{}, err
	}
	request.Header = headers
	return request, nil
}

func Do(method string, url string, token string, body interface{}) (*Resp, error) {
	client := newClient()
	req, err := newRequest(token, method, url, body)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do http request: %w", err)
	}

	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}
	var respResult Resp
	err = json.Unmarshal(result, &respResult)
	if err != nil {
		return nil, fmt.Errorf("cannot read resp body: %w", err)
	}

	if respResult.Code == 0 && resp.StatusCode == http.StatusOK {
		// 请求成功
		return &respResult, nil
	}

	return nil, fmt.Errorf("request fail code: %d ,msg: %s", respResult.Code, respResult.Msg)
}
