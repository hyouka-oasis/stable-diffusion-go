package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type HttpClient struct {
	client         *http.Client
	defaultHeaders map[string]string
}

func NewHttpClient(second int, defaultHeaders map[string]string) *HttpClient {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // disable verify
	}
	if second == 0 {
		return &HttpClient{
			client:         &http.Client{Transport: transCfg},
			defaultHeaders: defaultHeaders,
		}
	}
	return &HttpClient{
		client:         &http.Client{Timeout: time.Duration(second) * time.Second, Transport: transCfg},
		defaultHeaders: defaultHeaders,
	}
}

func (h *HttpClient) GetReader(urlStr string, params map[string]string, headers map[string]string) (status int, reader io.ReadCloser, header http.Header, err error) {
	// 创建 URL 对象
	u, err := url.Parse(urlStr)
	if err != nil {
		return 0, nil, nil, err
	}

	// 添加查询参数
	query := u.Query()
	for key, value := range params {
		query.Set(key, value)
	}
	u.RawQuery = query.Encode()

	// 创建 GET 请求
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return 0, nil, nil, err
	}

	// 设置请求头部
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
	if h.defaultHeaders != nil {
		for key, value := range h.defaultHeaders {
			req.Header.Set(key, value)
		}
	}

	// 发送请求并获取响应
	resp, err := h.client.Do(req)
	if err != nil {
		return 0, nil, nil, err
	}
	// 检查响应状态码(小于400 都算正常返回)
	if resp.StatusCode >= http.StatusBadRequest {
		// 读取响应的内容
		body, err := io.ReadAll(resp.Body)
		fmt.Println(err, string(body), "错误信息1")
		return resp.StatusCode, nil, resp.Header, fmt.Errorf("get \" %s \" failed with status code: %d , body: %s", urlStr, resp.StatusCode, string(body))
	}

	return resp.StatusCode, resp.Body, resp.Header, nil
}
