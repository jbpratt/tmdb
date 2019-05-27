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

func (c *Client) request(uri url.URL, method string) (*http.Response, error) {
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

func (c *Client) constructURI(uri string, options map[string]string) (*url.URL, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("api_key", c.APIKey)

	if options != nil {
		for key, value := range options {
			q.Add(key, value)
		}
	}

	u.RawQuery = q.Encode()

	return u, nil
}

func HandleResponse(res *http.Response) (string, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	return string(body), nil
}
