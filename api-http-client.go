package apihttpclient

import (
	"encoding/json"
	"net/http"
	"strings"

	httptool "github.com/alex19pov31/http-tool"
)

func NewApiHTTPClient(baseURL string, client *http.Client) *ApiHTTPClient {
	return &ApiHTTPClient{
		baseURL: baseURL,
		client:  client,
	}
}

type ApiHTTPClient struct {
	baseURL string
	client  *http.Client
	headers map[string]string
	cookies []http.Cookie
}

func (a *ApiHTTPClient) SetCookies(cookies ...http.Cookie) *ApiHTTPClient {
	a.cookies = cookies
	return a
}

func (a *ApiHTTPClient) AddCookie(cookie *http.Cookie) *ApiHTTPClient {
	a.cookies = append(a.cookies, *cookie)
	return a
}

func (a *ApiHTTPClient) SetHeaders(headers map[string]string) *ApiHTTPClient {
	a.headers = headers
	return a
}

func (a *ApiHTTPClient) AddHeader(name, value string) *ApiHTTPClient {
	a.headers[name] = value
	return a
}

func (a *ApiHTTPClient) GetRequest(url string) *httptool.ResultRequest {
	return httptool.CustomHTTPRequest("GET", a.baseURL+url, []byte{}, a.client, a.updRequest)
}

func (a *ApiHTTPClient) PostRequest(url string, data []byte) *httptool.ResultRequest {
	return httptool.CustomHTTPRequest("POST", a.baseURL+url, data, a.client, a.updRequest)
}

func (a *ApiHTTPClient) PostJSONRequest(url string, data interface{}) *httptool.ResultRequest {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return &httptool.ResultRequest{
			Error: err,
		}
	}

	return a.PostRequest(url, jsonData)
}

func (a *ApiHTTPClient) PutRequest(url string, data []byte) *httptool.ResultRequest {
	return httptool.CustomHTTPRequest("PUT", url, data, a.client, a.updRequest)
}

func (a *ApiHTTPClient) PutJSONRequest(url string, data interface{}) *httptool.ResultRequest {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return &httptool.ResultRequest{
			Error: err,
		}
	}

	return a.PutRequest(url, jsonData)
}

func (a *ApiHTTPClient) DeleteRequest(url string) *httptool.ResultRequest {
	return httptool.CustomHTTPRequest("DELETE", a.baseURL+url, []byte{}, a.client, a.updRequest)
}

func (a *ApiHTTPClient) updRequest(req *http.Request) *http.Request {
	req = httptool.SetChromeHeaders(req)
	for key, value := range a.headers {
		req.Header.Add(key, value)
	}

	for _, cookie := range a.cookies {
		req.AddCookie(&cookie)
	}

	return req
}

func (a *ApiHTTPClient) CompileGetParams(filter ...string) string {
	if len(filter) == 0 {
		return ""
	}

	return "?" + strings.Join(filter, "&")
}
