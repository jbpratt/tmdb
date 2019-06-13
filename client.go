package tmdb

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const BaseURI = "https://api.themoviedb.org/3"

var (
	ErrNoQueryParam = errors.New("no query paramater was supplied")
)

type Client struct {
	Client *http.Client
	APIKey string
}

func New(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}

func (c *Client) request(uri url.URL, method string) (string, error) {
	req, err := http.NewRequest(method, uri.String(), nil)
	if err != nil {
		return "", err
	}

	if c.Client == nil {
		c.Client = &http.Client{
			Timeout: time.Second * 10,
		}
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		defer res.Body.Close()
		return "", fmt.Errorf("recieved status code %d", res.StatusCode)
	}
	s, err := handleResponse(res)
	if err != nil {
		return "", err
	}
	return s, nil
}

func (c *Client) constructURI(uri string) (*url.URL, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("api_key", c.APIKey)
	u.RawQuery = q.Encode()
	return u, nil
}

func handleResponse(res *http.Response) (string, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	return string(body), nil
}

func sumQuery(query []string) string {
	var out string
	for _, q := range query {
		out += q + " "
	}
	return out
}

func setQuery(u *url.URL, querySlice []string) *url.URL {
	query := sumQuery(querySlice)
	var options = map[string]string{
		"query": url.QueryEscape(query),
	}

	q := u.Query()
	if options != nil {
		for key, value := range options {
			q.Add(key, value)
		}
	}

	u.RawQuery = q.Encode()
	return u
}
