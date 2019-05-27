package tmdb

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const BaseURI = "https://api.themoviedb.org/3"

var (
	ErrNoParams = errors.New("no paramaters were supplied")
)

type Client struct {
	Client *http.Client
	APIKey string
}

func New(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}

func (c *Client) Request(uri url.URL, method string) (*http.Response, error) {
	req, err := http.NewRequest(method, uri.String(), nil)
	if err != nil {
		return nil, err
	}

	if c.Client == nil {
		c.Client = &http.Client{
			Timeout: time.Second * 10,
		}
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		defer res.Body.Close()
		return nil, fmt.Errorf("recieved status code %d", res.StatusCode)
	}
	return res, nil
}
