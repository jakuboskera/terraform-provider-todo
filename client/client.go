package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/jakuboskera/terraform-provider-todo/models"
)

type Client struct {
	url        string
	api_key    string
	insecure   bool
	httpClient *http.Client
}

// NewClient creates common settings
func NewClient(url string, api_key string, insecure bool) *Client {

	return &Client{
		url:        url,
		api_key:    api_key,
		insecure:   insecure,
		httpClient: &http.Client{},
	}
}

// SendRequest send a http request
func (c *Client) SendRequest(method string, path string, payload interface{}, statusCode int) (value string, respheaders string, respCode int, err error) {
	url := c.url + path
	client := &http.Client{}

	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(payload)
	if err != nil {
		return "", "", 0, err
	}

	if c.insecure == true {
		tr := &http.Transport{
			Proxy:           http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{Transport: tr}
	}

	req, err := http.NewRequest(method, url, b)
	if err != nil {
		return "", "", 0, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", c.api_key)

	resp, err := client.Do(req)
	if err != nil {
		return "", "", resp.StatusCode, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", resp.StatusCode, err
	}
	resp.Body.Close()

	strbody := string(body)

	respHeaders := resp.Header
	headers, err := json.Marshal(respHeaders)
	if err != nil {
		return "", "", resp.StatusCode, err
	}

	if statusCode != 0 {
		if resp.StatusCode != statusCode {

			return "", "", 0, fmt.Errorf("[ERROR] unexpected status code got: %v expected: %v \n %v", resp.StatusCode, statusCode, strbody)
		}
	}

	return strbody, string(headers), resp.StatusCode, nil
}

// GetID gets the resource id from location response header
func GetID(body string) (id string, err error) {
	var jsonData models.ResponseHeaders
	err = json.Unmarshal([]byte(body), &jsonData)
	if err != nil {
		return "", err
	}

	location := jsonData.Location[0]

	url, err := url.Parse(location)
	if err != nil {
		return "", err
	}

	location = strings.Replace(url.Path, "/api/v1", "", -1)

	return location, nil
}
