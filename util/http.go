package util

import (
"bytes"
"encoding/json"
"errors"
"io"
"io/ioutil"
"net/http"
"net/url"
"qudian.com/go-crm/src/library/utils/jwt"
"strings"
"time"
)

// Res json Response struct
type Res struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type HttpClient struct {
	JwtKey    string
	JwtSecret string
	JwtEnable bool
}

func (c *HttpClient) Get(urlStr string, params map[string]interface{}, httpHeaders map[string]interface{}) ([]byte, error) {
	if len(params) > 0 {
		values := url.Values{}
		for k, v := range params {
			values.Add(k, v.(string))
		}
		urlStr += "?" + values.Encode()
	}

	//LogPrintf("get request params %v", params)
	//LogPrintf("get request urlStr %s", urlStr)

	req, err := http.NewRequest("GET", urlStr, nil)
	if httpHeaders != nil {
		for key, value := range httpHeaders {
			req.Header.Set(key, value.(string))
		}
	}
	if err != nil {
		//LogPrintf("create request error %s", err)
	}
	return c.doRequest(req, false)
}

func (c *HttpClient) GetWithFullData(urlStr string, params map[string]interface{}, httpHeaders map[string]interface{}) ([]byte, error) {
	if len(params) > 0 {
		values := url.Values{}
		for k, v := range params {
			values.Add(k, v.(string))
		}
		urlStr += "?" + strings.ReplaceAll(values.Encode(), "+", "%20")
	}

	//LogPrintf("get request params %v", params)
	//LogPrintf("get request urlStr %s", urlStr)

	req, err := http.NewRequest("GET", urlStr, nil)
	if httpHeaders != nil {
		for key, value := range httpHeaders {
			req.Header.Set(key, value.(string))
		}
	}
	if err != nil {
		//LogPrintf("create request error %s", err)
	}
	return c.doRequest(req, true)
}

func (c *HttpClient) Post(url string, params map[string]interface{}, httpHeaders map[string]interface{}) ([]byte, error) {
	var body io.Reader
	if params != nil {
		query, err := json.Marshal(params)
		if err != nil {
			panic(err)
		}
		body = bytes.NewBuffer(query)
	}

	//LogPrintf("post request params %v", params)
	//LogPrintf("post request url %s", url)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	if httpHeaders != nil {
		for key, value := range httpHeaders {
			req.Header.Set(key, value.(string))
		}
	}
	return c.doRequest(req, false)
}

func (c *HttpClient) PostWithFullData(url string, params interface{}, httpHeaders map[string]interface{}) ([]byte, error) {
	var body io.Reader
	if params != nil {
		query, err := json.Marshal(params)
		if err != nil {
			panic(err)
		}
		body = bytes.NewBuffer(query)
	}

	//LogPrintf("post request params %v", params)
	//LogPrintf("post request url %s", url)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	if httpHeaders != nil {
		for key, value := range httpHeaders {
			req.Header.Set(key, value.(string))
		}
	}
	return c.doRequest(req, true)
}

func (c *HttpClient) doRequest(req *http.Request, returnFullData bool) ([]byte, error) {
	// default request timeout is 3 second
	client := &http.Client{Timeout: 10 * time.Second}

	if c.JwtEnable {
		tokenString := jwt.GenerateJwtToken(c.JwtKey, c.JwtSecret)
		req.Header.Set("Authorization", "Bearer "+tokenString)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	//LogPrintf("request result %s", string(body))

	jsonUtil := JsonUtil{}
	code := jsonUtil.GetInt(body, 0, "code")
	if code != 200 && code != 0 {
		err := errors.New(jsonUtil.GetString(body, "message"))
		return []byte{}, err
	}
	if returnFullData {
		return body, nil
	} else {
		return jsonUtil.Get(body, "data"), nil
	}
}

func (c *HttpClient) GetOri(urlStr string, params map[string]interface{}, httpHeaders map[string]interface{}) ([]byte, error) {
	if len(params) > 0 {
		values := url.Values{}
		for k, v := range params {
			values.Add(k, v.(string))
		}
		urlStr += "?" + values.Encode()
	}

	//LogPrintf("get request params %v", params)
	//LogPrintf("get request urlStr %s", urlStr)

	req, err := http.NewRequest("GET", urlStr, nil)
	if httpHeaders != nil {
		for key, value := range httpHeaders {
			req.Header.Set(key, value.(string))
		}
	}
	if err != nil {
		//LogPrintf("create request error %s", err)
	}

	// default request timeout is 3 second
	client := &http.Client{Timeout: 10 * time.Second}

	if c.JwtEnable {
		tokenString := jwt.GenerateJwtToken(c.JwtKey, c.JwtSecret)
		req.Header.Set("Authorization", "Bearer "+tokenString)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	//LogPrintf("request result %s", string(body))

	return body, nil
}

func (c *HttpClient) PostOri(urlStr string, params map[string]interface{}, httpHeaders map[string]interface{}) ([]byte, error) {
	if len(params) > 0 {
		values := url.Values{}
		for k, v := range params {
			values.Add(k, v.(string))
		}
		urlStr += "?" + values.Encode()
	}

	//LogPrintf("post request params %v", params)
	//LogPrintf("post request urlStr %s", urlStr)

	req, err := http.NewRequest("POST", urlStr, nil)
	if httpHeaders != nil {
		for key, value := range httpHeaders {
			req.Header.Set(key, value.(string))
		}
	}
	if err != nil {
		//LogPrintf("create request error %s", err)
	}

	// default request timeout is 3 second
	client := &http.Client{Timeout: 10 * time.Second}

	if c.JwtEnable {
		tokenString := jwt.GenerateJwtToken(c.JwtKey, c.JwtSecret)
		req.Header.Set("Authorization", "Bearer "+tokenString)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	//LogPrintf("request result %s", string(body))

	return body, nil
}

