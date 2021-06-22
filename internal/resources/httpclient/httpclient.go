package httpclient

import (
	"bulbasaur/internal/config"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type response struct {
	Code   int
	Header http.Header
	Body   []byte
}

type Client struct {
	httpClient *http.Client
	baseUrl    string
	reqHeaders http.Header
	response   *response
}

var defaultHeaders = http.Header{}

func New(cfg *config.Resource) (client *Client, err error) {
	httpClient := http.Client{Timeout: 15 * time.Second}
	client = &Client{httpClient: &httpClient}
	defaultHeaders.Set("Content-Type", "application/json")

	for key, val := range cfg.Options {
		switch key {
		case "base_url":
			if strings.TrimSpace(val) == "" {
				err = errors.New("invalid parameter `base_url`")
				return
			}

			client.baseUrl = val
		case "timeout":
			var timeout time.Duration
			timeout, err = time.ParseDuration(val)
			if err != nil {
				return
			}
			client.httpClient.Timeout = timeout
		case "headers":
			for _, header := range strings.Split(val, ",") {
				str := strings.Split(header, "=")
				if len(str) != 2 {
					err = errors.New("invalid parameter headers, must be: `[key1]=[value1],[key2]=[value2]`")
					return
				}
				defaultHeaders.Set(str[0], str[1])
			}
		default:
			err = errors.New(fmt.Sprintf("invalid parameter: %s", key))
		}
	}

	client.reqHeaders = defaultHeaders
	return
}

func (c *Client) Check() error {
	resp, err := c.httpClient.Get(c.baseUrl)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusServiceUnavailable {
		return errors.New("http/client: server unavailable > " + c.baseUrl)
	}
	return nil
}

func (c *Client) Reset() error {
	c.response = nil
	c.reqHeaders = defaultHeaders
	return nil
}

func (c *Client) Close() error {
	return c.Close()
}

func (c *Client) SetRequestHeader(key, value string) error {
	c.reqHeaders.Set(key, value)
	return nil
}

func (c *Client) Request(method, path string, body []byte) (err error) {
	req, err := http.NewRequest(method, path, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	req.Header = c.reqHeaders
	baseURL, err := url.Parse(c.baseUrl)
	if err != nil {
		return
	}
	req.URL.Scheme = baseURL.Scheme
	req.URL.Host = baseURL.Host

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	c.response = &response{resp.StatusCode, resp.Header, body}
	return
}

func (c *Client) GetResponse() (statusCode int, headers http.Header, body []byte, err error) {
	if c.response != nil {
		statusCode = c.response.Code
		headers = c.response.Header
		body = c.response.Body
	}
	return
}
