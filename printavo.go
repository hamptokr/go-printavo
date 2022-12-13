// Package printavo provides an http client to the printavo API
package printavo

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	baseURL        = "https://www.printavo.com/"
	apiVersionPath = "api/v1/"
	userAgent      = "go-printavo"
)

// A Client manages the communication with the Printavo API
type Client struct {
	// HTTP client used to communicate with the API
	client *http.Client
	// Base URL for API requests. Defaults to the public API url but can be
	// overridden. Base URL should always contain a trailing slash
	baseURL *url.URL
	// Will be used in authentication. These are sent in the request as query
	// params
	email, token string
	UserAgent    string

	// Services
	AccountService *AccountService
	OrdersService  *OrdersService
}

type ListMetaData struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	TotalCount int `json:"total_count"`
	TotalPages int `json:"total_pages"`
}

type ListResponse[M any] struct {
	Data []M          `json:"data"`
	Meta ListMetaData `json:"meta"`
}

type ListOptions struct {
	Page    int `json:"page,omitempty"`
	PerPage int `json:"per_page,omitempty"`
}

func NewClient(email string, token string) (*Client, error) {
	c := &Client{UserAgent: userAgent}

	c.setBaseURL(baseURL)

	c.client = &http.Client{}
	c.email = email
	c.token = token

	c.AccountService = &AccountService{client: c}
	c.OrdersService = &OrdersService{client: c}
	return c, nil
}

func (c *Client) BaseURL() *url.URL {
	u := *c.baseURL
	return &u
}

func (c *Client) setBaseURL(urlStr string) error {
	// Guarantees the url has a trailing slash
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	if !strings.HasSuffix(baseURL.Path, apiVersionPath) {
		baseURL.Path += apiVersionPath
	}

	c.baseURL = baseURL

	return nil
}

func (c *Client) NewRequest(method, path string, opt interface{}) (*http.Request, error) {
	u := *c.baseURL
	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}

	u.RawPath = c.baseURL.Path + path
	u.Path = c.baseURL.Path + unescaped

	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")

	if c.UserAgent != "" {
		reqHeaders.Set("User-Agent", c.UserAgent)
	}

	body := new(bytes.Buffer)

	// If we are using POST or PUT, encode opt as the json body. Otherwise we
	// encode it as query params.
	switch {
	case method == http.MethodPost || method == http.MethodPut:
		reqHeaders.Set("Content-Type", "application/json")

		if opt != nil {
			jsonBody, err := json.Marshal(opt)
			if err != nil {
				return nil, err
			}

			_, err = body.Write(jsonBody)
			if err != nil {
				return nil, err
			}
		}
	case opt != nil:
		q, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	for k, v := range reqHeaders {
		req.Header[k] = v
	}

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	q := req.URL.Query()
	q.Add("email", c.email)
	q.Add("token", c.token)

	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return resp, err
}
